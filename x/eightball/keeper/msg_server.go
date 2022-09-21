package keeper

import (
	"context"
	"time"

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

	fortuneList, _ := k.GetAllFortunes(ctx)

	// check if there is already a fortune belonging to the msg sender, if there is reject the tx
	for _, fortune := range fortuneList.Fortunes {
		if fortune.Owner == msg.Sender {
			return nil, types.ErrAlreadyFortunate
		}
	}

	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return nil, err
	}

	// send offering to eightball module account to be transferred over to simple-dex
	// TODO: make offering non-nullable
	if err := k.bankKeeper.SendCoinsFromAccountToModule(ctx, sender, types.ModuleName, sdk.NewCoins(msg.Offering)); err != nil {
		return nil, err
	}

	// send an IBC token transfer to the addr owned by the eightball module account on the simple-dex host chain

	// get simple-dex chain connection ID & port ID to get the ICA account addr below
	dexConnectionID, found := k.GetDexConnectionID(ctx)
	if !found {
		return nil, types.ErrDexConnectionNotFound
	}

	eightballAddr := k.accountKeeper.GetModuleAddress(types.ModuleName)

	portID, err := icatypes.NewControllerPortID(eightballAddr.String())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "could not find account: %s", err)
	}

	// TODO: store ICA address so you don't have to recreate
	eightballICAAddr, found := k.icacontrollerKeeper.GetInterchainAccountAddress(ctx, dexConnectionID, portID)
	if !found {
		return nil, status.Errorf(codes.NotFound, "no account found for portID %s", portID)
	}

	// collect the channel id, and ica address to send an IBC transfer from the eightball module account to the ica account
	// it controls on host chain simple-dex
	dexChannelId, found := k.GetDexTransferChannelID(ctx)
	if !found {
		return nil, types.ErrDexConnectionNotFound
	}

	// grab the next sequence in the channel which will be the sequence number of this transfer packet
	sequence, found := k.ibcKeeper.ChannelKeeper.GetNextSequenceSend(ctx, k.transferKeeper.GetPort(ctx), dexChannelId)
	if !found {
		return nil, sdkerrors.Wrapf(
			channeltypes.ErrSequenceSendNotFound,
			"source port: %s, source channel: %s", k.transferKeeper.GetPort(ctx), dexChannelId,
		)
	}

	// create the request workflow

	//  associate the packet send with the workflow so that we can continue workflow
	// on packet ack
	workflow := types.NewWorkflow(sender, sdk.Coin{})
	k.SetPacketToWorkflow(ctx, types.SrcOrigin, transfertypes.PortID, dexChannelId, sequence, workflow)

	// send an IBC transfer from the eightball module account to the ica account
	// it controls on host chain simple-dex
	err = k.transferKeeper.SendTransfer(
		ctx,
		k.transferKeeper.GetPort(ctx),
		dexChannelId,
		msg.Offering,
		eightballAddr,
		eightballICAAddr,
		clienttypes.ZeroHeight(),
		uint64(ctx.BlockTime().Add(6*time.Hour).UnixNano()),
	)

	if err != nil {
		return nil, err
	}

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

	res, err := k.ibcKeeper.ChannelOpenInit(goCtx, channOpenMsg)

	if err != nil {
		return nil, err
	}

	k.SetDexTransferChannelID(ctx, res.ChannelId)

	if err := k.icacontrollerKeeper.RegisterInterchainAccount(ctx, msg.ConnectionId, eightballAddr.String()); err != nil {
		return nil, err
	}

	return &types.MsgConnectToDexResponse{}, nil
}
