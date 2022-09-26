package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/eightball module sentinel errors
var (
	ErrInvalidPacketTimeout  = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion        = sdkerrors.Register(ModuleName, 3, "invalid version")
	ErrDexConnectionNotFound = sdkerrors.Register(ModuleName, 4, "dex connection id not found")
	ErrDexChannelNotFound    = sdkerrors.Register(ModuleName, 5, "dex channel id not found")
	ErrOffererNotFound       = sdkerrors.Register(ModuleName, 6, "original offerer not found")
)
