package types

import (
	errorsmod "cosmossdk.io/errors"
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgCreateEventPost{}

func NewMsgCreateEventPost(creator string, title string, eventType string, datetime string, venue string, description string, interest string) *MsgCreateEventPost {
	return &MsgCreateEventPost{
		Creator:     creator,
		Title:       title,
		EventType:   eventType,
		Datetime:    datetime,
		Venue:       venue,
		Description: description,
		Interest:    interest,
	}
}

func (msg *MsgCreateEventPost) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Creator)
	if err != nil {
		return errorsmod.Wrapf(sdkerrors.ErrInvalidAddress, "invalid creator address (%s)", err)
	}
	return nil
}
