package keeper

import (
	"errors"
	"fmt"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
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

func (k Keeper) MintFortune(ctx sdk.Context, data transfertypes.FungibleTokenPacketData, offerer sdk.AccAddress) (types.Fortune, error) {
	fortuneList, _ := k.GetUnownedFortunes(ctx)
	var availableFortunes []types.Fortune
	var err error

	for _, fortune := range fortuneList.Fortunes {
		if fortune.Owner == "" {
			availableFortunes = append(availableFortunes, fortune)
		}
	}

	transferAmount, _ := sdk.NewIntFromString(data.Amount)
	transferCoin := sdk.NewCoin(data.Denom, transferAmount)

	// if there are no available fortunes, refund the tokens back to the original offerer
	if len(availableFortunes) == 0 {
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, offerer, sdk.NewCoins(transferCoin)); err != nil {
			return types.Fortune{}, err
		}
		return types.Fortune{}, fmt.Errorf("available fortunes list is empty. original FortuneList: %#v", fortuneList)
	}

	// use block time here to generate random index to enforce determinism
	randInt := int(ctx.BlockTime().UnixNano()) % len(availableFortunes)
	selectedFortune := availableFortunes[randInt]

	// check to see if the amount that has been transferred is enough to mint a fortune, if not, error
	fortunePrice, err := sdk.ParseCoinNormalized(selectedFortune.Price)
	if err != nil {
		return types.Fortune{}, err
	}

	if transferCoin.IsLT(fortunePrice) {
		return types.Fortune{}, errors.New("not enough money offered, try again")
	}

	selectedFortune.Owner = offerer.String()
	k.SetFortune(ctx, selectedFortune)

	// refund the rest of the tokens to original sender
	leftoverCoin := transferCoin.SubAmount(fortunePrice.Amount)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, offerer, sdk.NewCoins(leftoverCoin))

	return selectedFortune, nil
}
