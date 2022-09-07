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
	ibckeeper "github.com/cosmos/ibc-go/v3/modules/core/keeper"
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

		ibcKeeper           *ibckeeper.Keeper
		transferKeeper      *transferkeeper.Keeper
		scopedKeeper        capabilitykeeper.ScopedKeeper
		icacontrollerKeeper *icacontrollerkeeper.Keeper

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

	ibcKeeper *ibckeeper.Keeper,
	transferKeeper *transferkeeper.Keeper,
	scopedKeeper capabilitykeeper.ScopedKeeper,
	icacontrollerKeeper *icacontrollerkeeper.Keeper,
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

		ibcKeeper:           ibcKeeper,
		transferKeeper:      transferKeeper,
		scopedKeeper:        scopedKeeper,
		icacontrollerKeeper: icacontrollerKeeper,
		msgRouter:           msgRouter,
	}
}

func (k Keeper) Logger(ctx sdk.Context) log.Logger {
	return ctx.Logger().With("module", fmt.Sprintf("x/%s", types.ModuleName))
}

func (k Keeper) SetDexTransferChannelID(ctx sdk.Context, channelID string) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyDexTransferChannel(types.ModuleName), []byte(channelID))
}

func (k Keeper) GetDexTransferChannelID(ctx sdk.Context) (string, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyDexTransferChannel(types.ModuleName)

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

func (k Keeper) SetTransferSeqToOfferer(ctx sdk.Context, transferSeq uint64, offerer sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyTransferSeq(transferSeq), offerer)
}

func (k Keeper) GetTransferSeqToOfferer(ctx sdk.Context, transferSeq uint64) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyTransferSeq(transferSeq)
	if !store.Has(key) {
		return nil, false
	}

	return store.Get(key), true
}

func (k Keeper) SetICASeqToOfferer(ctx sdk.Context, icaSeq uint64, offerer sdk.AccAddress) {
	store := ctx.KVStore(k.storeKey)
	store.Set(types.KeyICASeq(icaSeq), offerer)
}

func (k Keeper) GetICASeqToOfferer(ctx sdk.Context, icaSeq uint64) (sdk.AccAddress, bool) {
	store := ctx.KVStore(k.storeKey)
	key := types.KeyICASeq(icaSeq)
	if !store.Has(key) {
		return nil, false
	}

	return store.Get(key), true
}
