package keeper

import (
	"context"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Fortunes(c context.Context, req *types.QueryFortunesRequest) (*types.QueryFortunesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fortunes []types.Fortune
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fortunesStore := prefix.NewStore(store, types.KeyPrefix(types.FortunesKeyPrefix))

	pageRes, err := query.Paginate(fortunesStore, req.Pagination, func(key []byte, value []byte) error {
		var fortune types.Fortune
		if err := k.cdc.Unmarshal(value, &fortune); err != nil {
			return err
		}

		fortunes = append(fortunes, fortune)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryFortunesResponse{Fortunes: fortunes, Pagination: pageRes}, nil
}

func (k Keeper) Fortune(c context.Context, req *types.QueryFortuneRequest) (*types.QueryFortuneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFortunes(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFortuneResponse{Fortune: val}, nil
}
