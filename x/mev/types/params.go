package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	sdkerrors "github.com/cosmos/cosmos-sdk/types/errors"
	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Custom parameter store error codes
const (
	CodeInvalidParam uint32 = 1
)

var (
	KeyMinBundleFee  = []byte("MinBundleFee")
	KeyMaxBundleSize = []byte("MaxBundleSize")
)

// Parameter store keys
var (
	ParamStoreKeyMinBundleFee  = []byte("MinBundleFee")
	ParamStoreKeyMaxBundleSize = []byte("MaxBundleSize")
)

// Params defines the parameters for the MEV module
type Params struct {
	MinBundleFee  sdk.Coin `json:"min_bundle_fee" yaml:"min_bundle_fee"`
	MaxBundleSize uint64   `json:"max_bundle_size" yaml:"max_bundle_size"`
}

// ParamKeyTable for MEV module
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// DefaultParams returns default parameters
func DefaultParams() Params {
	return Params{
		MinBundleFee:  sdk.NewCoin("usei", sdk.NewInt(1000)),
		MaxBundleSize: 100,
	}
}

// ParamSetPairs implements params.ParamSet
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyMinBundleFee, &p.MinBundleFee, validateMinBundleFee),
		paramtypes.NewParamSetPair(ParamStoreKeyMaxBundleSize, &p.MaxBundleSize, validateMaxBundleSize),
	}
}

func validateMinBundleFee(i interface{}) error {
	v, ok := i.(sdk.Coin)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	if !v.IsValid() {
		return sdkerrors.Wrap(sdkerrors.ErrInvalidCoins, "invalid coin")
	}

	return nil
}

func validateMaxBundleSize(i interface{}) error {
	_, ok := i.(uint64)
	if !ok {
		return sdkerrors.Wrapf(sdkerrors.ErrInvalidType, "invalid parameter type: %T", i)
	}

	return nil
}

// Validate validates parameters
func (p Params) Validate() error {
	if err := validateMinBundleFee(p.MinBundleFee); err != nil {
		return err
	}
	if err := validateMaxBundleSize(p.MaxBundleSize); err != nil {
		return err
	}
	return nil
}
