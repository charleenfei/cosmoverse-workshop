package keeper

import sdk "github.com/cosmos/cosmos-sdk/types"

func (k Keeper) OnTransferChannelOpen(ctx sdk.Context, channelID string) error {
	return nil
}

func (k Keeper) OnICAChannelOpen(ctx sdk.Context, channelID string) error {
	return nil
}
