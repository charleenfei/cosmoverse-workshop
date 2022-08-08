package keeper

import (
	"context"

	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	"github.com/cosmos/cosmos-sdk/store/prefix"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) FortunesAll(c context.Context, req *types.QueryAllFortunesRequest) (*types.QueryAllFortunesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	var fortuness []types.Fortunes
	ctx := sdk.UnwrapSDKContext(c)

	store := ctx.KVStore(k.storeKey)
	fortunesStore := prefix.NewStore(store, types.KeyPrefix(types.FortunesKeyPrefix))

	pageRes, err := query.Paginate(fortunesStore, req.Pagination, func(key []byte, value []byte) error {
		var fortunes types.Fortunes
		if err := k.cdc.Unmarshal(value, &fortunes); err != nil {
			return err
		}

		fortuness = append(fortuness, fortunes)
		return nil
	})

	if err != nil {
		return nil, status.Error(codes.Internal, err.Error())
	}

	return &types.QueryAllFortunesResponse{Fortunes: fortuness, Pagination: pageRes}, nil
}

func (k Keeper) Fortunes(c context.Context, req *types.QueryGetFortunesRequest) (*types.QueryGetFortunesResponse, error) {
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

	return &types.QueryGetFortunesResponse{Fortunes: val}, nil
}
