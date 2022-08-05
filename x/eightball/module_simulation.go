package eightball

import (
	"math/rand"

	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/sample"
	eightballsimulation "github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/simulation"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	simappparams "github.com/cosmos/cosmos-sdk/simapp/params"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
	"github.com/cosmos/cosmos-sdk/x/simulation"
)

// avoid unused import issue
var (
	_ = sample.AccAddress
	_ = eightballsimulation.FindAccount
	_ = simappparams.StakePerAccount
	_ = simulation.MsgEntryKind
	_ = baseapp.Paramspace
)

const (
	opWeightMsgFeelingLucky = "op_weight_msg_feeling_lucky"
	// TODO: Determine the simulation weight value
	defaultWeightMsgFeelingLucky int = 100

	// this line is used by starport scaffolding # simapp/module/const
)

// GenerateGenesisState creates a randomized GenState of the module
func (AppModule) GenerateGenesisState(simState *module.SimulationState) {
	accs := make([]string, len(simState.Accounts))
	for i, acc := range simState.Accounts {
		accs[i] = acc.Address.String()
	}
	eightballGenesis := types.GenesisState{
		Params: types.DefaultParams(),
		// this line is used by starport scaffolding # simapp/module/genesisState
	}
	simState.GenState[types.ModuleName] = simState.Cdc.MustMarshalJSON(&eightballGenesis)
}

// ProposalContents doesn't return any content functions for governance proposals
func (AppModule) ProposalContents(_ module.SimulationState) []simtypes.WeightedProposalContent {
	return nil
}

// RandomizedParams creates randomized  param changes for the simulator
func (am AppModule) RandomizedParams(_ *rand.Rand) []simtypes.ParamChange {

	return []simtypes.ParamChange{}
}

// RegisterStoreDecoder registers a decoder
func (am AppModule) RegisterStoreDecoder(_ sdk.StoreDecoderRegistry) {}

// WeightedOperations returns the all the gov module operations with their respective weights.
func (am AppModule) WeightedOperations(simState module.SimulationState) []simtypes.WeightedOperation {
	operations := make([]simtypes.WeightedOperation, 0)

	var weightMsgFeelingLucky int
	simState.AppParams.GetOrGenerate(simState.Cdc, opWeightMsgFeelingLucky, &weightMsgFeelingLucky, nil,
		func(_ *rand.Rand) {
			weightMsgFeelingLucky = defaultWeightMsgFeelingLucky
		},
	)
	operations = append(operations, simulation.NewWeightedOperation(
		weightMsgFeelingLucky,
		eightballsimulation.SimulateMsgFeelingLucky(am.accountKeeper, am.bankKeeper, am.keeper),
	))

	// this line is used by starport scaffolding # simapp/module/operation

	return operations
}
