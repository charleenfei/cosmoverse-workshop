package eightball

import (
	"errors"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	channeltypes "github.com/cosmos/ibc-go/v6/modules/core/04-channel/types"
	porttypes "github.com/cosmos/ibc-go/v6/modules/core/05-port/types"
	ibcexported "github.com/cosmos/ibc-go/v6/modules/core/exported"

	icatypes "github.com/cosmos/ibc-go/v6/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
)

var _ porttypes.Middleware = &IBCMiddleware{}

// IBCMiddleware implements the ICS26 callbacks for the fee middleware given the
// fee keeper and the underlying application.
type IBCMiddleware struct {
	app    porttypes.IBCModule
	keeper keeper.Keeper
}

// NewIBCMiddleware creates a new IBCMiddlware given the keeper and underlying application
func NewIBCMiddleware(app porttypes.IBCModule, k keeper.Keeper) IBCMiddleware {
	return IBCMiddleware{
		app:    app,
		keeper: k,
	}
}

// OnChanOpenInit implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenInit(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID string,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	version string,
) (string, error) {
	// call underlying callback
	return im.app.OnChanOpenInit(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, version)
}

// OnChanOpenTry implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenTry(
	ctx sdk.Context,
	order channeltypes.Order,
	connectionHops []string,
	portID,
	channelID string,
	chanCap *capabilitytypes.Capability,
	counterparty channeltypes.Counterparty,
	counterpartyVersion string,
) (string, error) {
	// call underlying callback
	return im.app.OnChanOpenTry(ctx, order, connectionHops, portID, channelID, chanCap, counterparty, counterpartyVersion)
}

// OnChanOpenAck implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenAck(
	ctx sdk.Context,
	portID,
	channelID,
	counterpartyChannelID,
	counterpartyVersion string,
) error {
	switch portID {
	case transfertypes.PortID:
		im.keeper.OnTransferChannelOpen(ctx, channelID)
	default:
		im.keeper.OnICAChannelOpen(ctx, channelID)
	}
	// call underlying callback
	return im.app.OnChanOpenAck(ctx, portID, channelID, counterpartyChannelID, counterpartyVersion)
}

// OnChanOpenConfirm implements the IBCModule interface
func (im IBCMiddleware) OnChanOpenConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	switch portID {
	case transfertypes.PortID:
		im.keeper.OnTransferChannelOpen(ctx, channelID)
	case icatypes.PortID:
		im.keeper.OnICAChannelOpen(ctx, channelID)
	default:
		return errors.New("invalid port")
	}
	// call underlying callback
	return im.app.OnChanOpenConfirm(ctx, portID, channelID)
}

// OnChanCloseInit implements the IBCModule interface
func (im IBCMiddleware) OnChanCloseInit(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// call underlying callback
	return im.app.OnChanCloseInit(ctx, portID, channelID)
}

// OnChanCloseConfirm implements the IBCModule interface
func (im IBCMiddleware) OnChanCloseConfirm(
	ctx sdk.Context,
	portID,
	channelID string,
) error {
	// call underlying callback
	return im.app.OnChanCloseConfirm(ctx, portID, channelID)
}

// OnRecvPacket implements the IBCModule interface
func (im IBCMiddleware) OnRecvPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) ibcexported.Acknowledgement {
	// call underlying callback
	ack := im.app.OnRecvPacket(ctx, packet, relayer)
	if ack.Success() {
		var transferData transfertypes.FungibleTokenPacketData
		if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &transferData); err == nil {
			workflow, found := im.keeper.GetPacketToWorkflow(ctx, types.DstOrigin, packet.DestinationPort, packet.DestinationChannel, packet.Sequence)
			// the transfer recv packet arrived before the acknowledgement
			// we can't associate this packet with the offerer until we get the ICA Ack
			// but we can save the swapped coins here to retrieve later
			transferAmount, ok := sdk.NewIntFromString(transferData.Amount)
			if !ok {
				err := sdkerrors.Wrapf(transfertypes.ErrInvalidAmount, "unable to parse transfer amount (%s) into math.Int", transferData.Amount)
				return channeltypes.NewErrorAcknowledgement(err)
			}
			transferCoin := sdk.NewCoin(transferData.Denom, transferAmount)

			if !found {

				workflow := types.NewWorkflow(sdk.AccAddress{}, transferCoin)
				im.keeper.SetPacketToWorkflow(ctx, types.DstOrigin, packet.DestinationPort, packet.DestinationChannel, packet.Sequence, workflow)
			} else {
				workflow.SwappedCoin = transferCoin
				_, err = im.keeper.MintFortune(ctx, workflow)
				if err != nil {
					return channeltypes.NewErrorAcknowledgement(err)
				}
			}

		}

	}
	return ack
}

// OnAcknowledgementPacket implements the IBCModule interface
func (im IBCMiddleware) OnAcknowledgementPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	acknowledgement []byte,
	relayer sdk.AccAddress,
) error {
	// both transfer and ICA use the default acknowledgement
	var ack channeltypes.Acknowledgement
	if err := channeltypes.SubModuleCdc.UnmarshalJSON(acknowledgement, &ack); err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrUnknownRequest, "cannot unmarshal default channel acknowledgement: %v", err)
	}

	var transferData transfertypes.FungibleTokenPacketData
	var icaData icatypes.InterchainAccountPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &transferData); err == nil {
		err = im.keeper.OnTransferAck(ctx, transferData, packet, ack.Success())
		if err != nil {
			return sdkerrors.Wrap(err, "failed eightball transfer ack callback")
		}
	} else if err := icatypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &icaData); err == nil {
		err = im.keeper.OnICAAck(ctx, icaData, packet, ack)
		if err != nil {
			return sdkerrors.Wrap(err, "failed eightball ica ack callback")
		}
	} else {
		return errors.New("packet data invalid")
	}

	// call underlying callback
	return im.app.OnAcknowledgementPacket(ctx, packet, acknowledgement, relayer)
}

// OnTimeoutPacket implements the IBCModule interface
func (im IBCMiddleware) OnTimeoutPacket(
	ctx sdk.Context,
	packet channeltypes.Packet,
	relayer sdk.AccAddress,
) error {
	var transferData transfertypes.FungibleTokenPacketData
	var icaData icatypes.InterchainAccountPacketData
	if err := transfertypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &transferData); err == nil {
		im.keeper.OnTransferTimeout(ctx, transferData)
	} else if err := icatypes.ModuleCdc.UnmarshalJSON(packet.GetData(), &icaData); err == nil {
		im.keeper.OnICATimeout(ctx, icaData)
	} else {
		return errors.New("packet data invalid")
	}

	// call underlying callback
	return im.app.OnTimeoutPacket(ctx, packet, relayer)
}

// SendPacket implements the ICS4 Wrapper interface
func (im IBCMiddleware) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	// call underlying callback
	return im.keeper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (im IBCMiddleware) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	// call underlying callback
	return im.keeper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}

func (im IBCMiddleware) GetAppVersion(ctx sdk.Context, portID, channelID string) (string, bool) {
	return im.keeper.GetAppVersion(ctx, portID, channelID)
}
