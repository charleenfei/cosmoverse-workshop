package keeper

import (
	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

// GetAllFortunes returns all fortunes
func (k Keeper) GetAllFortunes(ctx sdk.Context) (list []types.Fortune) {
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
