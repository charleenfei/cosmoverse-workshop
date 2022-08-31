package types

// DONTCOVER

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

// x/eightball module sentinel errors
var (
	ErrAlreadyFortunate              = sdkerrors.Register(ModuleName, 1, "you already got your fortune!")
	ErrInvalidPacketTimeout = sdkerrors.Register(ModuleName, 2, "invalid packet timeout")
	ErrInvalidVersion       = sdkerrors.Register(ModuleName, 3, "invalid version")
)
