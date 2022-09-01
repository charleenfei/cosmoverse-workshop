package keeper_test

// import (
// 	"testing"

// 	testkeeper "github.com/charleenfei/cosmoverse-workshop/testutil/keeper"
// 	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
// 	sdk "github.com/cosmos/cosmos-sdk/types"
// 	"github.com/stretchr/testify/require"
// )

// func TestParamsQuery(t *testing.T) {
// 	keeper, ctx := testkeeper.EightballKeeper(t)
// 	wctx := sdk.WrapSDKContext(ctx)
// 	params := types.DefaultParams()
// 	keeper.SetParams(ctx, params)

// 	response, err := keeper.Params(wctx, &types.QueryParamsRequest{})
// 	require.NoError(t, err)
// 	require.Equal(t, &types.QueryParamsResponse{Params: params}, response)
// }
