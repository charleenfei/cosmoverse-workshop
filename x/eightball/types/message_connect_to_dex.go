package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

const TypeMsgConnectToDex = "connect_to_dex"

var _ sdk.Msg = &MsgConnectToDex{}

func NewMsgConnectToDex(creator string, connectionId string) *MsgConnectToDex {
  return &MsgConnectToDex{
		Creator: creator,
    ConnectionId: connectionId,
	}
}

func (msg *MsgConnectToDex) Route() string {
  return RouterKey
}

func (msg *MsgConnectToDex) Type() string {
  return TypeMsgConnectToDex
}

func (msg *MsgConnectToDex) GetSigners() []sdk.AccAddress {
  creator, err := sdk.AccAddressFromBech32(msg.Creator)
  if err != nil {
    panic(err)
  }
  return []sdk.AccAddress{creator}
}

func (msg *MsgConnectToDex) GetSignBytes() []byte {
  bz := ModuleCdc.MustMarshalJSON(msg)
  return sdk.MustSortJSON(bz)
}

func (msg *MsgConnectToDex) ValidateBasic() error {
  _, err := sdk.AccAddressFromBech32(msg.Creator)
  	if err != nil {
  		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
  	}
  return nil
}

