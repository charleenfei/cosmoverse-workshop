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

	// check if there is already a fortune belonging to the msg sender, if there is reject the tx
	for _, fortune := range fortunes {
		if fortune.Owner == msg.Creator {
			return nil, types.ErrAlreadyFortunate
		}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return nil, err
	}

	// send offering to eightball module account to be transferred over to simple-dex
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(*msg.Offering)); err != nil {
		return nil, err
	}

	// send an IBC token transfer to the addr owned by the eightball module account on the simple-dex host chain
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
