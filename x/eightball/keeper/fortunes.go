package keeper

import (
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
)

// SetFortune set a specific fortune in the store from its index
func (k Keeper) SetFortune(ctx sdk.Context, fortunes types.Fortune) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))
	b := k.cdc.MustMarshal(&fortunes)
	store.Set(types.FortuneKey(
		fortunes.Owner,
	), b)
}

// GetFortunes returns a fortune from its index
func (k Keeper) GetFortune(
	ctx sdk.Context,
	owner string,

) (val types.Fortune, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))

	b := store.Get(types.FortuneKey(
		owner,
	))
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

func (k Keeper) SetUnownedFortunes(ctx sdk.Context, fortunes []types.Fortune) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))
	fortuneList := types.FortuneList{
		Fortunes: fortunes,
	}
	b := types.ModuleCdc.MustMarshal(&fortuneList)
	store.Set(types.UnownedFortuneKey(), b)
}

func (k Keeper) GetUnownedFortunes(
	ctx sdk.Context,

) (val types.FortuneList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))

	b := store.Get(types.UnownedFortuneKey())
	if b == nil {
		return val, false
	}

	k.cdc.MustUnmarshal(b, &val)
	return val, true
}

// RemoveFortunes removes a fortunes from the store
func (k Keeper) RemoveFortunes(
	ctx sdk.Context,
	owner string,

) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))
	store.Delete(types.FortuneKey(
		owner,
	))
}

// GetAllFortunes returns all owned fortunes
func (k Keeper) GetAllOwnedFortunes(ctx sdk.Context) (list []types.Fortune) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fortune
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}

func (k Keeper) MintFortune(ctx sdk.Context, data icatypes.InterchainAccountPacketData, icaPacketSequence uint64) (types.Fortune, error) {
	fortuneList, _ := k.GetUnownedFortunes(ctx)
	var availableFortunes []types.Fortune

	for _, fortune := range fortuneList.Fortunes {
		if fortune.Owner == "" {
			availableFortunes = append(availableFortunes, fortune)
		}
	}

	// use block time here to generate random index to enforce determinism
	randInt := int(ctx.BlockTime().UnixNano()) % len(availableFortunes)
	selectedFortune := availableFortunes[randInt]

	initialOfferer, found := k.GetICASeqToOfferer(ctx, icaPacketSequence)
	if !found {
		return types.Fortune{}, types.ErrOffererNotFound
	}

	selectedFortune.Owner = initialOfferer.String()
	k.SetFortune(ctx, selectedFortune)

	return selectedFortune, nil
}
