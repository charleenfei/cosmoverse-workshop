package keeper

import (
	"fmt"
	"math/rand"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"

	transfertypes "github.com/cosmos/ibc-go/v6/modules/apps/transfer/types"
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

func (k Keeper) SetAllFortunes(ctx sdk.Context, fortunes []types.Fortune) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))
	fortuneList := types.FortuneList{
		Fortunes: fortunes,
	}
	b := types.ModuleCdc.MustMarshal(&fortuneList)
	store.Set(types.AllFortuneKey(), b)
}

func (k Keeper) GetAllFortunes(
	ctx sdk.Context,

) (val types.FortuneList, found bool) {
	store := prefix.NewStore(ctx.KVStore(k.storeKey), types.KeyPrefix(types.FortuneKeyPrefix))

	b := store.Get(types.AllFortuneKey())
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

// called when workflow is complete
func (k Keeper) MintFortune(ctx sdk.Context, workflow types.Workflow) (types.Fortune, error) {
	fortuneList, _ := k.GetAllFortunes(ctx)
	var availableFortunes []types.Fortune
	var err error

	for _, fortune := range fortuneList.Fortunes {
		if fortune.Owner == "" {
			availableFortunes = append(availableFortunes, fortune)
		}
	}

	offerer := sdk.MustAccAddressFromBech32(workflow.Offerer)
	swappedCoin := workflow.SwappedCoin

	dexChannelId, found := k.GetDexTransferChannelID(ctx)
	if !found {
		return types.Fortune{}, fmt.Errorf("dexCHannelid not found")
	}

	ibcTrace := GetIBCTrace(swappedCoin.Denom, transfertypes.PortID, dexChannelId)
	ibcCoin := sdk.NewCoin(ibcTrace.IBCDenom(), swappedCoin.Amount)

	// if there are no available fortunes, refund the tokens back to the original offerer
	if len(availableFortunes) == 0 {
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, offerer, sdk.NewCoins(ibcCoin)); err != nil {
			return types.Fortune{}, err
		}
		return types.Fortune{}, fmt.Errorf("available fortunes list is empty. original FortuneList: %#v", fortuneList)
	}

	// use block time here to generate random index to enforce determinism
	rand.Seed(ctx.BlockTime().UnixNano())
	index := rand.Intn(len(availableFortunes))
	selectedFortune := availableFortunes[index]

	// check to see if the amount that has been transferred is enough to mint a fortune, if not, error
	fortunePrice, err := sdk.ParseCoinNormalized(selectedFortune.Price)
	if err != nil {
		return types.Fortune{}, err
	}

	if swappedCoin.IsLT(fortunePrice) {
		if err = k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, offerer, sdk.NewCoins(ibcCoin)); err != nil {
			return types.Fortune{}, err
		}
		return types.Fortune{}, fmt.Errorf("available fortunes list is empty. original FortuneList: %#v", fortuneList)
	}

	selectedFortune.Owner = offerer.String()
	k.SetFortune(ctx, selectedFortune)

	fortuneList.Fortunes[index] = selectedFortune
	k.SetAllFortunes(ctx, fortuneList.Fortunes)

	// refund the rest of the tokens to original sender
	leftoverCoin := ibcCoin.SubAmount(fortunePrice.Amount)
	k.bankKeeper.SendCoinsFromModuleToAccount(ctx, types.ModuleName, offerer, sdk.NewCoins(leftoverCoin))

	return selectedFortune, nil
}

// GetIBCTrace returns the trace of the full token denom sent to the receiving channel
func GetIBCTrace(fullTokenDenom string, portID, channelID string) transfertypes.DenomTrace {
	return transfertypes.ParseDenomTrace(fmt.Sprintf("%s/%s/%s", portID, channelID, fullTokenDenom))
}
