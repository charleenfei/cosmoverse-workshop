package keeper

import (
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	host "github.com/cosmos/ibc-go/v3/modules/core/24-host"
	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"
)

func (k Keeper) OnTransferAck(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData, ackSuccess bool) error {
	if ackSuccess {

		portID, err := icatypes.NewControllerPortID(transferData.Sender)
		if err != nil {
			return err
		}

			// getconnectionID from connect to dex , stub out for now
		channelID, found := k.icacontrollerKeeper.GetActiveChannelID(ctx, ConnectionId, portID)
		if !found {
			sdkerrors.Wrapf(icatypes.ErrActiveChannelNotFound, "failed to retrieve active channel for port %s", portID)
		}

		chanCap, found := k.scopedKeeper.GetCapability(ctx, host.ChannelCapabilityPath(portID, channelID))
		if !found {
			sdkerrors.Wrap(channeltypes.ErrChannelCapabilityNotFound, "module does not own channel capability")
		}

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
		_, err = k.icacontrollerKeeper.SendTx(ctx, chanCap, msg.ConnectionId, portID, packetData, uint64(timeoutTimestamp))
		if err != nil {
			return err
		}
	}

	return nil
}

func (k Keeper) OnICAAck(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData, ackSuccess bool) error {
	// if ackSuccess {
	// 	k.
	// }

	return nil
}

func (k Keeper) OnTransferTimeout(ctx sdk.Context, transferData transfertypes.FungibleTokenPacketData) error {

	return nil
}

func (k Keeper) OnICATimeout(ctx sdk.Context, icaData icatypes.InterchainAccountPacketData) error {

	return nil
}
