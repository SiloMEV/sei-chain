package types

import (
	"fmt"

	paramtypes "github.com/cosmos/cosmos-sdk/x/params/types"
)

// Parameter store keys
var (
	ParamStoreKeyEnabled = []byte("Enabled")
	// Add other parameter keys as needed
)

// ParamKeyTable returns the parameter key table.
func ParamKeyTable() paramtypes.KeyTable {
	return paramtypes.NewKeyTable().RegisterParamSet(&Params{})
}

// NewParams creates a new Params instance
func NewParams(enabled bool) Params {
	return Params{
		Enabled: enabled,
	}
}

// DefaultParams returns default mev module parameters
func DefaultParams() Params {
	return Params{
		Enabled: true,
	}
}

// Validate performs basic validation on mev parameters.
func (p Params) Validate() error {
	// Add parameter validation as needed
	return nil
}

// ParamSetPairs returns the parameter set pairs.
func (p *Params) ParamSetPairs() paramtypes.ParamSetPairs {
	return paramtypes.ParamSetPairs{
		paramtypes.NewParamSetPair(ParamStoreKeyEnabled, &p.Enabled, validateBool),
	}
}

func validateBool(i interface{}) error {
	_, ok := i.(bool)
	if !ok {
		return fmt.Errorf("invalid parameter type: %T", i)
	}

	return nil
}
