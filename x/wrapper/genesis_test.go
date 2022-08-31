package wrapper_test

import (
	"testing"

	keepertest "github.com/charleenfei/cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/cosmoverse-workshop/testutil/nullify"
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper"
	"github.com/charleenfei/cosmoverse-workshop/x/wrapper/types"
	"github.com/stretchr/testify/require"
)

func TestGenesis(t *testing.T) {
	genesisState := types.GenesisState{
		Params: types.DefaultParams(),
		PortId: types.PortID,
		// this line is used by starport scaffolding # genesis/test/state
	}

	k, ctx := keepertest.WrapperKeeper(t)
	wrapper.InitGenesis(ctx, *k, genesisState)
	got := wrapper.ExportGenesis(ctx, *k)
	require.NotNil(t, got)

	nullify.Fill(&genesisState)
	nullify.Fill(got)

	require.Equal(t, genesisState.PortId, got.PortId)

	// this line is used by starport scaffolding # genesis/test/assert
}
