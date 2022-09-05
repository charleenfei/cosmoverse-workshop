package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgFeelingLucky = "feeling_lucky"

var _ sdk.Msg = &MsgFeelingLucky{}

func NewMsgFeelingLucky(sender string, offering *sdk.Coin) *MsgFeelingLucky {
	return &MsgFeelingLucky{
		Sender:   sender,
		Offering: offering,
	}
}

func (msg *MsgFeelingLucky) Route() string {
	return RouterKey
}

func (msg *MsgFeelingLucky) Type() string {
	return TypeMsgFeelingLucky
}

func (msg *MsgFeelingLucky) GetSigners() []sdk.AccAddress {
	creator, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{creator}
}

func (msg *MsgFeelingLucky) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

func (msg *MsgFeelingLucky) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
