package types

import (
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
)

// DefaultGenesis returns default genesis state as raw bytes for the mev
// module.
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		Bundles: []Bundle{},
	}
}

// NewGenesisState creates a new genesis state.
func NewGenesisState(params Params, bundles []Bundle) *GenesisState {
	return &GenesisState{
		Params:  params,
		Bundles: bundles,
	}
}

// ValidateGenesis performs genesis state validation for the mev module.
func ValidateGenesis(data GenesisState) error {
	if err := data.Params.Validate(); err != nil {
		return err
	}

	for _, bundle := range data.Bundles {
		if err := ValidateBundle(bundle); err != nil {
			return err
		}
	}
	return nil
}

// GetGenesisStateFromAppState returns x/mev GenesisState given raw application
// genesis state.
func GetGenesisStateFromAppState(cdc codec.JSONCodec, appState map[string]json.RawMessage) *GenesisState {
	var genesisState GenesisState

	if appState[ModuleName] != nil {
		cdc.MustUnmarshalJSON(appState[ModuleName], &genesisState)
	}

	return &genesisState
}
