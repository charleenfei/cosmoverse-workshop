package types

import sdk "github.com/cosmos/cosmos-sdk/types"

func NewWorkflow(offerer sdk.AccAddress, coin sdk.Coin) Workflow {
	return Workflow{
		Offerer:     offerer.String(),
		SwappedCoin: coin,
	}
}
