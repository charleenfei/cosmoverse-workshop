package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/charleenfei/cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.WrapperKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
