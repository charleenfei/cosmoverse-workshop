package keeper

import (
	"context"

	"github.com/charleenfei/cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"

	icatypes "github.com/cosmos/ibc-go/v3/modules/apps/27-interchain-accounts/types"
	transfertypes "github.com/cosmos/ibc-go/v3/modules/apps/transfer/types"
	clienttypes "github.com/cosmos/ibc-go/v3/modules/core/02-client/types"
	channeltypes "github.com/cosmos/ibc-go/v3/modules/core/04-channel/types"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
		if fortune.Owner == msg.Sender {
			return nil, types.ErrAlreadyFortunate
		}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// send offering to eightball module account to be transferred over to simple-dex
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(*msg.Offering)); err != nil {
		return nil, err
	}

	// send an IBC token transfer to the addr owned by the eightball module account on the simple-dex host chain
	eightballAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	// use connection ID to get the ICA account addr below
	dexConnectionID, found := k.GetDexConnectionID(ctx)
	if !found {
		return nil, types.ErrDexConnectionNotFound
	}

	dexChannelId, found := k.GetDexChannelID(ctx)
	if !found {
		return nil, types.ErrDexConnectionNotFound
	}

	portID, err := icatypes.NewControllerPortID(eightballAddr.String())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	eightballICAAddr, found := k.icacontrollerKeeper.GetInterchainAccountAddress(ctx, dexConnectionID, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	k.transferKeeper.SendTransfer(
		ctx,
		k.transferKeeper.GetPort(ctx),
		dexChannelId,
		*msg.Offering,
		eightballAddr,
		eightballICAAddr,
		clienttypes.ZeroHeight(),
		0,
	)

	return &types.MsgFeelingLuckyResponse{}, nil
}

func (k msgServer) ConnectToDex(goCtx context.Context, msg *types.MsgConnectToDex) (*types.MsgConnectToDexResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	eightballAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)
	k.SetDexConnectionID(ctx, msg.ConnectionId)

	channOpenMsg := channeltypes.NewMsgChannelOpenInit(
		k.transferKeeper.GetPort(ctx),
		transfertypes.Version,
		channeltypes.UNORDERED,
		[]string{msg.ConnectionId},
		transfertypes.PortID,
		eightballAddr.String(),
	)

	// TODO: is this correct?
	handler := k.msgRouter.Handler(msg)

	res, err := handler(ctx, channOpenMsg)
	if err != nil {
		return nil, err
	}

	ctx.EventManager().EmitEvents(res.GetEvents())
	firstMsgResponse := res.MsgResponses[0]
	channelOpenInitResponse, ok := firstMsgResponse.GetCachedValue().(*channeltypes.MsgChannelOpenInitResponse)

	if !ok {
		return nil, sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "failed to covert %T message response to %T", firstMsgResponse.GetCachedValue(), &channeltypes.MsgChannelOpenInitResponse{})
	}

	k.SetDexChannelID(ctx, channelOpenInitResponse.ChannelId)

	if err := k.icacontrollerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, eightballAddr.String()); err != nil {
		return nil, err
	}

	return &types.MsgConnectToDexResponse{}, nil
}
