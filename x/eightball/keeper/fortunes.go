package keeper

import (
	"math/rand"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"

	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	"google.golang.org/protobuf/proto"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"

	simpledextypes "github.com/charleenfei/simple-dex/types"
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

func (k Keeper) MintFortune(ctx sdk.Context, data icatypes.InterchainAccountPacketData) (types.Fortune, error) {
	fortunes := k.GetAllFortunes(ctx)
	var availableFortunes []types.Fortune

	for _, fortune := range fortunes {
		if fortune.Owner == "" {
			availableFortunes = append(availableFortunes, fortune)
		}
	}

	selectedFortune := availableFortunes[rand.Intn(len(availableFortunes)-1)]

	// TODO: simple dex -> public to import types
	msgResponse := &simpledextypes.MsgSwapResponse{}
	if err := proto.Unmarshal(data.Data, msgResponse); err != nil {
		return types.Fortune{}, sdkerrors.Wrapf(sdkerrors.ErrJSONUnmarshal, "cannot unmarshal delegate response message: %s", err.Error())
	}

	// TODO: set owner and price
	selectedFortune.Owner = getSenderFromDatafromMsgResponse
	k.SetFortune(ctx, selectedFortune)

	return selectedFortune, nil
}
