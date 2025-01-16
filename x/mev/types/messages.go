package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var _ sdk.Msg = &MsgSubmitBundle{}

// Message types for the MEV module
const (
	TypeMsgSubmitBundle = "submit_bundle"
)

// NewMsgSubmitBundle creates a new MsgSubmitBundle instance
func NewMsgSubmitBundle(bundle Bundle) *MsgSubmitBundle {
	return &MsgSubmitBundle{
		Sender:    bundle.Sender,
		Txs:       bundle.Txs,
		BlockNum:  bundle.BlockNum,
		Timestamp: bundle.Timestamp,
	}
}

// Route implements the sdk.Msg interface
func (msg MsgSubmitBundle) Route() string {
	return RouterKey
}

// Type implements the sdk.Msg interface
func (msg MsgSubmitBundle) Type() string {
	return TypeMsgSubmitBundle
}

// GetSigners implements the sdk.Msg interface
func (msg MsgSubmitBundle) GetSigners() []sdk.AccAddress {
	sender, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		panic(err)
	}
	return []sdk.AccAddress{sender}
}

// GetSignBytes implements the sdk.Msg interface
func (msg MsgSubmitBundle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(&msg)
	return sdk.MustSortJSON(bz)
}

// ValidateBasic implements the sdk.Msg interface
func (msg MsgSubmitBundle) ValidateBasic() error {
	_, err := sdk.AccAddressFromBech32(msg.Sender)
	if err != nil {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidAddress, "invalid sender address (%s)", err)
	}

	if len(msg.Txs) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "transactions cannot be empty")
	}

	if msg.BlockNum == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "block number must be greater than 0")
	}

	return nil
}
