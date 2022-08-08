package keeper_test

import (
	"strconv"
	"testing"

	keepertest "github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/nullify"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/keeper"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func createNFortunes(keeper *keeper.Keeper, ctx sdk.Context, n int) []types.Fortunes {
	items := make([]types.Fortunes, n)
	for i := range items {
		items[i].Owner = strconv.Itoa(i)

		keeper.SetFortunes(ctx, items[i])
	}
	return items
}

func TestFortunesGet(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNFortunes(keeper, ctx, 10)
	for _, item := range items {
		rst, found := keeper.GetFortunes(ctx,
			item.Owner,
		)
		require.True(t, found)
		require.Equal(t,
			nullify.Fill(&item),
			nullify.Fill(&rst),
		)
	}
}
func TestFortunesRemove(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNFortunes(keeper, ctx, 10)
	for _, item := range items {
		keeper.RemoveFortunes(ctx,
			item.Owner,
		)
		_, found := keeper.GetFortunes(ctx,
			item.Owner,
		)
		require.False(t, found)
	}
}

func TestFortunesGetAll(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	items := createNFortunes(keeper, ctx, 10)
	require.ElementsMatch(t,
		nullify.Fill(items),
		nullify.Fill(keeper.GetAllFortunes(ctx)),
	)
}
