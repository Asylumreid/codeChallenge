package keeper

import (
	"context"

	"crud/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) ExpressInterest(goCtx context.Context, msg *types.MsgExpressInterest) (*types.MsgExpressInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	// TODO: Handling the message
	_ = ctx

	return &types.MsgExpressInterestResponse{}, nil
}
