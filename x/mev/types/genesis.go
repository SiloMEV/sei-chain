package types

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params:  DefaultParams(),
		Bundles: []Bundle{},
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
