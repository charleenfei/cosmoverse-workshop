package keeper

import (
	"context"

	"github.com/charleenfei/icq-ics20-cosmoverse-workshop/x/eightball/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
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

	// Osmo as host chain
	// 8ball is the controller

	// parse into sdk.Coin
	// check if the offering corresponds to a fortune
	// if not send back rejected
	// if so
	// escrow funds (send from owner to module account)


	// module account send everything over ICS20 to ICA on DEX chain
		// OnAck:
			// if success: MultiMsgTx
				// send ICA message to swap user tokens to desired token denom
				// ICS20 send from ICA account on DEX chain desired denom to module account 
					// module acc --> community pool
				// ICS20 send from ICA account on DEX chain rest back to user
			// OnAck for MultiMsgTx:
				// send fortune to user
			// OnTimeout
			// OnAckErr
				//
			
		// OnTimeout: refund tokens back to user
		// OnAckErr: refund tokens


	// chain 2: DEX chain w LP
	// accept X amount, gives Y amount back

	_ = ctx

	return &types.MsgFeelingLuckyResponse{}, nil
}