package keeper

import (
	"context"
	"fmt"

	errorsmod "cosmossdk.io/errors"
	"crud/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

func (k msgServer) UpdateEventPost(goCtx context.Context, msg *types.MsgUpdateEventPost) (*types.MsgUpdateEventPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Post{
		Title:       msg.Title,
		EventType:   msg.EventType,
		Datetime:    msg.Datetime,
		Venue:       msg.Venue,
		Description: msg.Description,
		Creator:     msg.Creator,
		Id:          msg.Id,
	}
	val, found := k.GetPost(ctx, msg.Id)
	if !found {
		return nil, errorsmod.Wrap(sdkerrors.ErrKeyNotFound, fmt.Sprintf("key %d doesn't exist", msg.Id))
	}
	if msg.Creator != val.Creator {
		return nil, errorsmod.Wrap(sdkerrors.ErrUnauthorized, "incorrect owner")
	}
	k.SetPost(ctx, post)

	return &types.MsgUpdateEventPostResponse{}, nil
}
