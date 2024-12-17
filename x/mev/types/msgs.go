package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	proto "github.com/gogo/protobuf/proto"
)

// Message types for the MEV module
const (
	TypeMsgSubmitBundle = "submit_bundle"
)

// MsgSubmitBundle defines a message to submit a bundle of transactions
type MsgSubmitBundle struct {
	Sender       sdk.AccAddress `json:"sender" yaml:"sender"`
	Transactions [][]byte       `json:"transactions" yaml:"transactions"`
	BlockHeight  int64          `json:"block_height" yaml:"block_height"`
	BundleFee    sdk.Coin       `json:"bundle_fee" yaml:"bundle_fee"`
}

// Implement the proto.Message interface
func (msg *MsgSubmitBundle) Reset()         { *msg = MsgSubmitBundle{} }
func (msg *MsgSubmitBundle) String() string { return proto.CompactTextString(msg) }
func (msg *MsgSubmitBundle) ProtoMessage()  {}

// Route implements sdk.Msg
func (msg *MsgSubmitBundle) Route() string {
	return RouterKey
}

// Type implements sdk.Msg
func (msg *MsgSubmitBundle) Type() string {
	return TypeMsgSubmitBundle
}

// ValidateBasic implements sdk.Msg
func (msg *MsgSubmitBundle) ValidateBasic() error {
	if msg.Sender.Empty() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidAddress, "sender address cannot be empty")
	}

	if len(msg.Transactions) == 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "transactions cannot be empty")
	}

	if len(msg.Transactions) > 100 { // or whatever max size you want to enforce
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "too many transactions in bundle")
	}

	if msg.BlockHeight <= 0 {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidRequest, "block height must be positive")
	}

	if !msg.BundleFee.IsValid() || msg.BundleFee.Amount.IsZero() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid bundle fee")
	}

	return nil
}

// GetSignBytes implements sdk.Msg
func (msg *MsgSubmitBundle) GetSignBytes() []byte {
	bz := ModuleCdc.MustMarshalJSON(msg)
	return sdk.MustSortJSON(bz)
}

// GetSigners implements sdk.Msg
func (msg *MsgSubmitBundle) GetSigners() []sdk.AccAddress {
	return []sdk.AccAddress{msg.Sender}
}

// Update the NewMsgSubmitBundle function
func NewMsgSubmitBundle(sender sdk.AccAddress, txs [][]byte, height int64, fee sdk.Coin) *MsgSubmitBundle {
	return &MsgSubmitBundle{
		Sender:       sender,
		Transactions: txs,
		BlockHeight:  height,
		BundleFee:    fee,
	}
}
