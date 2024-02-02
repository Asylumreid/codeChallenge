package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"crud/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) ExpressInterest(goCtx context.Context, msg *types.MsgExpressInterest) (*types.MsgExpressInterestResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	err := k.AddInterestToEventPost(ctx, msg.Id)
	if err != nil {
		return nil, err
	}

	post, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}

	return &types.MsgExpressInterestResponse{
		NewInterestCount: post.Interest,
	}, nil

}
