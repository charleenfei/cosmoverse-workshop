package eightball_test

import (
	"testing"

	keepertest "github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/nullify"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),

		FortunesList: []types.Fortunes{
			{
				Owner: "0",
			},
			{
				Owner: "1",
			},
		},
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.EightballKeeper(t)
	eightball.InitGenesis(ctx, *k, genesisState)
	got := eightball.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.ElementsMatch(t, genesisState.FortunesList, got.FortunesList)
	// this line is used by starport scaffolding # genesis/test/assert
}
