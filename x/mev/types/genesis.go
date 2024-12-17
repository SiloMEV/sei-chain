package types

// GenesisState defines the MEV module's genesis state
type GenesisState struct {
	Params Params `json:"params" yaml:"params"`
}

// NewGenesisState creates a new GenesisState instance
func NewGenesisState(params Params) GenesisState {
	return GenesisState{
		Params: params,
	}
}

// DefaultGenesis returns default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	return gs.Params.Validate()
}
