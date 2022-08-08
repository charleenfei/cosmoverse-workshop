package keeper

import (
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// SetFortunes set a specific fortunes in the store from its index
func (k Keeper) SetFortunes(ctx sdk.Context, fortunes types.Fortunes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortunesKeyPrefix))
	b := k.cdc.MustMarshal(&fortunes)
	store.Set(types.FortunesKey(
		fortunes.Owner,
	), b)
}

// GetFortunes returns a fortunes from its index
func (k Keeper) GetFortunes(
	ctx sdk.Context,
	owner string,

) (val types.Fortunes, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortunesKeyPrefix))

	b := store.Get(types.FortunesKey(
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
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortunesKeyPrefix))
	store.Delete(types.FortunesKey(
		owner,
	))
}

// GetAllFortunes returns all fortunes
func (k Keeper) GetAllFortunes(ctx sdk.Context) (list []types.Fortunes) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortunesKeyPrefix))
	iterator := sdk.KVStorePrefixIterator(store, []byte{})

	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var val types.Fortunes
		k.cdc.MustUnmarshal(iterator.Value(), &val)
		list = append(list, val)
	}

	return
}
