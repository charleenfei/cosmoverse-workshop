package keeper_test

import (
	"strconv"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/query"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"

	keepertest "github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/keeper"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/testutil/nullify"
	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
)

// Prevent strconv unused error
var _ = strconv.IntSize

func TestFortunesQuerySingle(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFortunes(keeper, ctx, 2)
	for _, tc := range []struct {
		desc     string
		request  *types.QueryGetFortunesRequest
		response *types.QueryGetFortunesResponse
		err      error
	}{
		{
			desc: "First",
			request: &types.QueryGetFortunesRequest{
				Owner: msgs[0].Owner,
			},
			response: &types.QueryGetFortunesResponse{Fortunes: msgs[0]},
		},
		{
			desc: "Second",
			request: &types.QueryGetFortunesRequest{
				Owner: msgs[1].Owner,
			},
			response: &types.QueryGetFortunesResponse{Fortunes: msgs[1]},
		},
		{
			desc: "KeyNotFound",
			request: &types.QueryGetFortunesRequest{
				Owner: strconv.Itoa(100000),
			},
			err: status.Error(codes.NotFound, "not found"),
		},
		{
			desc: "InvalidRequest",
			err:  status.Error(codes.InvalidArgument, "invalid request"),
		},
	} {
		t.Run(tc.desc, func(t *testing.T) {
			response, err := keeper.Fortunes(wctx, tc.request)
			if tc.err != nil {
				require.ErrorIs(t, err, tc.err)
			} else {
				require.NoError(t, err)
				require.Equal(t,
					nullify.Fill(tc.response),
					nullify.Fill(response),
				)
			}
		})
	}
}

func TestFortunesQueryPaginated(t *testing.T) {
	keeper, ctx := keepertest.EightballKeeper(t)
	wctx := sdk.WrapSDKContext(ctx)
	msgs := createNFortunes(keeper, ctx, 5)

	request := func(next []byte, offset, limit uint64, total bool) *types.QueryAllFortunesRequest {
		return &types.QueryAllFortunesRequest{
			Pagination: &query.PageRequest{
				Key:        next,
				Offset:     offset,
				Limit:      limit,
				CountTotal: total,
			},
		}
	}
	t.Run("ByOffset", func(t *testing.T) {
		step := 2
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FortunesAll(wctx, request(nil, uint64(i), uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Fortunes), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Fortunes),
			)
		}
	})
	t.Run("ByKey", func(t *testing.T) {
		step := 2
		var next []byte
		for i := 0; i < len(msgs); i += step {
			resp, err := keeper.FortunesAll(wctx, request(next, 0, uint64(step), false))
			require.NoError(t, err)
			require.LessOrEqual(t, len(resp.Fortunes), step)
			require.Subset(t,
				nullify.Fill(msgs),
				nullify.Fill(resp.Fortunes),
			)
			next = resp.Pagination.NextKey
		}
	})
	t.Run("Total", func(t *testing.T) {
		resp, err := keeper.FortunesAll(wctx, request(nil, 0, 0, true))
		require.NoError(t, err)
		require.Equal(t, len(msgs), int(resp.Pagination.Total))
		require.ElementsMatch(t,
			nullify.Fill(msgs),
			nullify.Fill(resp.Fortunes),
		)
	})
	t.Run("InvalidRequest", func(t *testing.T) {
		_, err := keeper.FortunesAll(wctx, nil)
		require.ErrorIs(t, err, status.Error(codes.InvalidArgument, "invalid request"))
	})
}
