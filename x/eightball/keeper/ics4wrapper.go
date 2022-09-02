package keeper

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitytypes "github.com/cosmos/cosmos-sdk/x/capability/types"
	ibcexported "github.com/cosmos/ibc-go/v3/modules/core/exported"
)

// SendPacket implements the ICS4 Wrapper interface
func (k Keeper) SendPacket(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
) error {
	// call underlying callback
	return k.ics4Wrapper.SendPacket(ctx, chanCap, packet)
}

// WriteAcknowledgement implements the ICS4 Wrapper interface
func (k Keeper) WriteAcknowledgement(
	ctx sdk.Context,
	chanCap *capabilitytypes.Capability,
	packet ibcexported.PacketI,
	ack ibcexported.Acknowledgement,
) error {
	// call underlying callback
	return k.ics4Wrapper.WriteAcknowledgement(ctx, chanCap, packet, ack)
}
