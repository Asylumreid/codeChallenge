package keeper

import (
	"context"

	"crud/x/crud/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

func (k msgServer) CreateEventPost(goCtx context.Context, msg *types.MsgCreateEventPost) (*types.MsgCreateEventPostResponse, error) {
	ctx := sdk.UnwrapSDKContext(goCtx)
	var post = types.Post{
		Title:       msg.Title,
		EventType:   msg.EventType,
		Datetime:    msg.Datetime,
		Venue:       msg.Venue,
		Description: msg.Description,
		Interest:    0,
		Creator:     msg.Creator,
	}
	id := k.AppendPost(
		ctx,
		post,
	)

	return &types.MsgCreateEventPostResponse{
		Id: id,
	}, nil
}
