package keeper

import (
	"context"

    "github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)


func (k msgServer) ConnectToDex(goCtx context.Context,  msg *types.MsgConnectToDex) (*types.MsgConnectToDexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

    // TODO: Handling the message
    _ = ctx

	return &types.MsgConnectToDexResponse{}, nil
}
