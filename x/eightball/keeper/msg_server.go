package keeper

import (
	"context"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"

	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
)

type msgServer struct {
	Keeper
}

// NewMsgServerImpl returns an implementation of the MsgServer interface
// for the provided Keeper.
func NewMsgServerImpl(keeper Keeper) types.MsgServer {
	return &msgServer{Keeper: keeper}
}

var _ types.MsgServer = msgServer{}

func (k msgServer) FeelingLucky(goCtx context.Context, msg *types.MsgFeelingLucky) (*types.MsgFeelingLuckyResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	fortunes := k.GetAllFortunes(ctx)

	for _, fortune := range fortunes {
		if fortune.Owner == msg.Creator {
			return nil, types.ErrAlreadyFortunate
		}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(*msg.Offering)); err != nil {
		return nil, err
	}

	eightballAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	k.transferKeeper.SendTransfer(
		ctx,
		"TODO",
		"TODO",
		*msg.Offering,
		eightballAddr,
		"TODO",
		clienttypes.NewHeight(1, 110),
		0,
	)

	return &types.MsgFeelingLuckyResponse{}, nil
}
