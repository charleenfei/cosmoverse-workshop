package keeper_test

import (
	"context"
	"testing"

	keepertest "github.com/charleenfei/cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func setupMsgServer(t testing.TB) (types.MsgServer, context.Context) {
	k, ctx := keepertest.EightballKeeper(t)
	return keeper.NewMsgServerImpl(*k), sdk.WrapSDKContext(ctx)
}
