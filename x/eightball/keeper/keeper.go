package keeper

import (
	"fmt"

	"github.com/tendermint/tendermint/libs/log"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"

	baseapp "github.com/cosmos/cosmos-sdk/baseapp"
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	capabilitykeeper "github.com/cosmos/cosmos-sdk/x/capability/keeper"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"

	icacontrollerkeeper "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/controller/keeper"
	transferkeeper "github.com/cosmos/ibc-go/v3/modules/apps/transfer/keeper"
	porttypes "github.com/cosmos/ibc-go/v3/modules/core/05-port/types"
)

type (
	Keeper struct {
		cdc        codec.BinaryCodec
		storeKey   sdk.StoreKey
		memKey     sdk.StoreKey
		paramstore paramtypes.Subspace

		accountKeeper types.AccountKeeper
		bankKeeper    types.BankKeeper

		ics4Wrapper porttypes.ICS4Wrapper

		transferKeeper      transferkeeper.Keeper
		scopedKeeper        capabilitykeeper.ScopedKeeper
		icacontrollerKeeper icacontrollerkeeper.Keeper

		msgRouter *baseapp.MsgServiceRouter
	}
)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey,
	memKey sdk.StoreKey,
	ps paramtypes.Subspace,

	accountKeeper types.AccountKeeper,
	bankKeeper types.BankKeeper,

	ics4Wrapper porttypes.ICS4Wrapper,

	transferKeeper transferkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	icacontrollerKeeper icacontrollerkeeper.Keeper,
	msgRouter *baseapp.MsgServiceRouter,
) *Keeper {
	// set KeyTable if it has not already been set
	if !ps.HasKeyTable() {
		ps = ps.WithKeyTable(types.ParamKeyTable())
	}

	return &Keeper{

		cdc:        cdc,
		storeKey:   storeKey,
		memKey:     memKey,
		paramstore: ps,

		accountKeeper: accountKeeper,
		bankKeeper:    bankKeeper,

		ics4Wrapper: ics4Wrapper,

		transferKeeper:      transferKeeper,
		scopedKeeper:        scopedKeeper,
		icacontrollerKeeper: icacontrollerKeeper,
		msgRouter:     msgRouter,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetDexChannelID(ctx sdk.Context, channelID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyDexChannel(types.ModuleName), []byte(channelID))
}

func (k Keeper) GetDexChannelID(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyDexChannel(types.ModuleName)

	if !store.Has(key) {
		return "", false
	}

	return string(store.Get(key)), true
}

func (k Keeper) SetDexConnectionID(ctx sdk.Context, connectionID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyDexConnection(types.ModuleName), []byte(connectionID))

}

func (k Keeper) GetDexConnectionID(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyDexConnection(types.ModuleName)
	if !store.Has(key) {
		return "", false
	}

	return string(store.Get(key)), true
}