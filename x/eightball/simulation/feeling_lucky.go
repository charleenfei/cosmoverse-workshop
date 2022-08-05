package simulation

import (
	"math/rand"

	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/keeper"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/baseapp"
	sdk "github.com/cosmos/cosmos-sdk/types"
	simtypes "github.com/cosmos/cosmos-sdk/types/simulation"
)

func SimulateMsgFeelingLucky(
	ak types.AccountKeeper,
	bk types.BankKeeper,
	k keeper.Keeper,
) simtypes.Operation {
	return func(r *rand.Rand, app *baseapp.BaseApp, ctx sdk.Context, accs []simtypes.Account, chainID string,
	) (simtypes.OperationMsg, []simtypes.FutureOperation, error) {
		simAccount, _ := simtypes.RandomAcc(r, accs)
		msg := &types.MsgFeelingLucky{
			Creator: simAccount.Address.String(),
		}

		// TODO: Handling the FeelingLucky simulation

		return simtypes.NoOpMsg(types.ModuleName, msg.Type(), "FeelingLucky simulation not implemented"), nil, nil
	}
}
