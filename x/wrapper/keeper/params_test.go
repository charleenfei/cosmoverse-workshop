package keeper_test

import (
	"testing"

	testkeeper "github.com/charleenfei/cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper/types"
	"github.com/stretchr/testify/require"
)

func TestGetParams(t *testing.T) {
	k, ctx := testkeeper.WrapperKeeper(t)
	params := types.DefaultParams()

	k.SetParams(ctx, params)

	require.EqualValues(t, params, k.GetParams(ctx))
}
