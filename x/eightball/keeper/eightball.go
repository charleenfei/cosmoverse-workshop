package keeper

import (
	"errors"
	"fmt"
	"time"

	proto "github.com/gogo/protobuf/proto"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	simpledextypes "github.com/charleenfei/simple-dex/simple-dex/x/simpledex/types"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
)

func (k Keeper) OnTransferAck(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData, packetSequence uint64, ackSuccess bool) error {
	if ackSuccess {

		// get simple-dex chain connection ID & port ID to get the ICA account addr below
		icaPortID, err := icatypes.NewControllerPortID(transferData.Sender)
		if err != nil {
			return err
		}

		dexConnectionID, found := k.GetDexConnectionID(ctx)
		if !found {
			return types.ErrDexConnectionNotFound
		}

		// TODO: store ICA addr to grab
		eightballICAAddr, found := k.icacontrollerKeeper.GetInterchainAccountAddress(ctx, dexConnectionID, icaPortID)
		if !found {
			return status.Errorf(codes.NotFound, "no account found for portID %s", icaPortID)
		}

		// get ica channel id to grant ica channel capability
		icaChannelID, found := k.icacontrollerKeeper.GetActiveChannelID(ctx, dexConnectionID, icaPortID)
		if !found {
			sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", icaPortID)
		}

		// TODO: this can be updated when ibc-go v6 is live
		chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(icaPortID, icaChannelID))
		if !found {
			sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
		}
		// get eightball module aaccount address to set as receiver of the transfer of exchanged tokens back from simple-dex
		eightballAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

		dexChannelId, found := k.GetDexTransferChannelID(ctx)
		if !found {
			return types.ErrDexChannelNotFound
		}

		// grab the dex channel so we can get the counterparty channel ID (as MsgSwap is executed on simple-dex chain, and sent over the channel on simple-dex back to us)
		channel, found := k.ibcKeeper.ChannelKeeper.GetChannel(ctx, transfertypes.PortID, dexChannelId)
		if !found {
			return types.ErrDexChannelNotFound
		}

		transferAmount, ok := sdk.NewIntFromString(transferData.Amount)
		if !ok {
			return sdkerrors.Wrapf(transfertypes.ErrInvalidAmount, "unable to parse transfer amount (%s) into math.Int", transferData.Amount)
		}

		// construct the denom that would have been created by counterparty
		// by creating trace and then hashing
		// the trace is created with destination port and channel identifiers
		trace := transfertypes.DenomTrace{
			Path:      fmt.Sprintf("%s/%s", transfertypes.PortID, channel.Counterparty.ChannelId),
			BaseDenom: transferData.Denom,
		}
		counterpartyDenom := trace.IBCDenom()

		// create the MsgSwap to be submitted by the ica controller account on simple-dex
		// set transfer port and dex channel id to tell simple-dex where to send back funds after they have been exchanged
		msgSwap := &simpledextypes.MsgSwap{
			Sender: eightballICAAddr,
			Offer: sdk.Coin{
				Denom:  counterpartyDenom,
				Amount: transferAmount,
			},
			MinAsk: sdk.Coin{
				Denom:  "token",
				Amount: sdk.NewInt(100),
			},
			PortId:    transfertypes.PortID,
			ChannelId: channel.Counterparty.ChannelId,
			Receiver:  eightballAddr.String(),
		}

		data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{msgSwap})
		if err != nil {
			return err
		}

		packetData := icatypes.InterchainAccountPacketData{
			Type: icatypes.EXECUTE_TX,
			Data: data,
		}

		// grab the next sequence in the channel which will be the sequence number of this ica packet
		sequence, found := k.ibcKeeper.ChannelKeeper.GetNextSequenceSend(ctx, icaPortID, icaChannelID)
		if !found {
			return sdkerrors.Wrapf(
				channeltypes.ErrSequenceSendNotFound,
				"source port: %s, source channel: %s", icaPortID, icaChannelID,
			)
		}

		// use the transfer sequence passed into this function to grab the sender of the FeelingLucky msg
		sender, found := k.GetTransferSeqToOfferer(ctx, packetSequence)
		if !found {
			return types.ErrOffererNotFound
		}

		//  save k/v pair which associates the sender of the FeelingLucky msg w the sequence number of this ica packet
		k.SetICASeqToOfferer(ctx, sequence, sender)

		// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
		timeoutTimestamp := ctx.BlockTime().Add(time.Hour).UnixNano()

		// send the packet data containing the MsgSwap to be executed on simple-dex chain
		_, err = k.icacontrollerKeeper.SendTx(ctx, chanCap, dexConnectionID, icaPortID, packetData, uint64(timeoutTimestamp))
		if err != nil {
			return err
		}
	}

	// TODO: bank send from module account back to sender transfer amt
	return nil
}

func (k Keeper) OnICAAck(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData, packetSequence uint64, ack channeltypes.Acknowledgement) error {
	if ack.Success() {
		// TODO: need some way to get receiver of transfer addr (owner of fortune) from simple dex & price that the owner paid
		// k.MintFortune(ctx, icaData, packetSequence)

		// TODO: handle extra token overflow that are sent from simple-dex
		txMsgData := &sdk.TxMsgData{}
		if err := proto.Unmarshal(ack.GetResult(), txMsgData); err != nil {
			return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal ICS-27 tx message data: %v", err)
		}

		switch len(txMsgData.Data) {
		case 1:
			var swapResponse simpledextypes.MsgSwapResponse
			if err := proto.Unmarshal(txMsgData.Data[0].Data, &swapResponse); err != nil {
				return err
			}

			offerer, found := k.GetICASeqToOfferer(ctx, packetSequence)
			if !found {
				return errors.New("ica seq not found")
			}

			k.SetTransferRecvSeqToOfferer(ctx, swapResponse.Sequence, offerer)

		default:
			return errors.New("unexpected number of messages")
		}
	}

	// TODO: if ICA fails, send another ICA message that transfers the amount back to sender

	return nil
}

func (k Keeper) OnTransferTimeout(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData) error {

	return nil
}

func (k Keeper) OnICATimeout(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData) error {

	return nil
}
