package keeper

import (
	"context"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (k Keeper) Fortunes(c context.Context, req *types.QueryFortunesRequest) (*types.QueryFortunesResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}

	ctx := sdk.UnwrapSDKContext(c)

	fortuneList, found := k.GetAllFortunes(ctx)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFortunesResponse{Fortunes: fortuneList.Fortunes}, nil
}

func (k Keeper) Fortune(c context.Context, req *types.QueryFortuneRequest) (*types.QueryFortuneResponse, error) {
	if req == nil {
		return nil, status.Error(codes.InvalidArgument, "invalid request")
	}
	ctx := sdk.UnwrapSDKContext(c)

	val, found := k.GetFortune(
		ctx,
		req.Owner,
	)
	if !found {
		return nil, status.Error(codes.NotFound, "not found")
	}

	return &types.QueryFortuneResponse{Fortune: val}, nil
}
