package keeper

import (
	"time"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	// simpledextypes "github.com/charleenfei/simple-dex/types"
)

func (k Keeper) OnTransferAck(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData, ackSuccess bool) error {
	if ackSuccess {

		portID, err := icatypes.NewControllerPortID(transferData.Sender)
		if err != nil {
			return err
		}

		dexConnectionID, found := k.GetDexConnectionID(ctx)
		if !found {
			return types.ErrDexConnectionNotFound
		}

		channelID, found := k.icacontrollerKeeper.GetActiveChannelID(ctx, dexConnectionID, portID)
		if !found {
			sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
		}

		// TODO: this can be updated when ibc-go v6 is live
		chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
		if !found {
			sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
		}

		// eightballICAAddr, found := k.icacontrollerKeeper.GetInterchainAccountAddress(ctx, dexConnectionID, portID)
		// if !found {
		// 	return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
		// }

		// msgSwap := &simpledextypes.MsgSwap{
		// 	Sender:
		// 	Offer:
		// 	MinAsk:
		// 	PortId:
		// 	ChannelId:
		// 	Receiver:
		// }

		data, err := icatypes.SerializeCosmosTx(k.cdc, []sdk.Msg{MsgSwap})
		if err != nil {
			return err
		}

		packetData := icatypes.InterchainAccountPacketData{
			Type: icatypes.EXECUTE_TX,
			Data: data,
		}

		// timeoutTimestamp set to max value with the unsigned bit shifted to sastisfy hermes timestamp conversion
		// it is the responsibility of the auth module developer to ensure an appropriate timeout timestamp
		timeoutTimestamp := ctx.BlockTime().Add(time.Minute).UnixNano()
		_, err = k.icacontrollerKeeper.SendTx(ctx, chanCap, dexConnectionID, portID, packetData, uint64(timeoutTimestamp))
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) OnICAAck(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData, ackSuccess bool) error {
	if ackSuccess {
		// TODO: need some way to get receiver of transfer addr (owner of fortune) from simple dex & price that the owner paid
		k.MintFortune(ctx, icaData)
	}

	return nil
}

func (k Keeper) OnTransferTimeout(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData) error {

	return nil
}

func (k Keeper) OnICATimeout(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData) error {

	return nil
}
