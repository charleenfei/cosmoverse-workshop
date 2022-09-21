package eightball

import (
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/keeper"
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// InitGenesis initializes the capability module's state from a provided genesis
// state.
func InitGenesis(ctx sdk.Context, k keeper.Keeper, genState types.GenesisState) {
	// Set all the fortune
	k.SetAllFortunes(ctx, genState.FortunesList)
	// this line is used by starport scaffolding # genesis/module/init
	k.SetParams(ctx, genState.Params)
}

// ExportGenesis returns the capability module's exported genesis.
func ExportGenesis(ctx sdk.Context, k keeper.Keeper) *types.GenesisState {
	genesis := types.DefaultGenesis()
	genesis.Params = k.GetParams(ctx)

	fortunesList, _ := k.GetAllFortunes(ctx)
	genesis.FortunesList = fortunesList.Fortunes
	// this line is used by starport scaffolding # genesis/module/export

	return genesis
}
