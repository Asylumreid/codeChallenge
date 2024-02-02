package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgUpdateEventPost{}

func NewMsgUpdateEventPost(creator string, title string, eventType string, datetime string, venue string, description string, interest string, id uint64) *MsgUpdateEventPost {
	return &MsgUpdateEventPost{
		Creator:     creator,
		Title:       title,
		EventType:   eventType,
		Datetime:    datetime,
		Venue:       venue,
		Description: description,
		Interest:    interest,
		Id:          id,
	}
}

func (msg *MsgUpdateEventPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
