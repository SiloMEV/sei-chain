package types

import (
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
)

var (
	ErrInvalidBundle      = sdkerrors.Register(ModuleName, 1, "invalid bundle")
	ErrBundleTooLarge     = sdkerrors.Register(ModuleName, 2, "bundle exceeds maximum size")
	ErrInsufficientFee    = sdkerrors.Register(ModuleName, 3, "insufficient bundle fee")
	ErrInvalidBlockHeight = sdkerrors.Register(ModuleName, 4, "invalid target block height")
	ErrBundleSimulation   = sdkerrors.Register(ModuleName, 5, "bundle simulation failed")
)
