package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"crud/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) DeleteEventPost(goCtx context.Context, msg *types.MsgDeleteEventPost) (*types.MsgDeleteEventPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)

	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.RemovePost(ctx, msg.Id)

	return &types.MsgDeleteEventPostResponse{}, nil
}
