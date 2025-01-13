package types

import (
	proto "github.com/gogo/protobuf/proto"
)

// GenesisState defines the mev module's genesis state
type GenesisState struct {
	// Bundles that should exist at genesis
	Bundles []Bundle `protobuf:"bytes,1,rep,name=bundles,proto3" json:"bundles,omitempty"`
}

func (m *GenesisState) Reset()         { *m = GenesisState{} }
func (m *GenesisState) String() string { return proto.CompactTextString(m) }
func (*GenesisState) ProtoMessage()    {}

// DefaultGenesisState returns the default genesis state
func DefaultGenesisState() *GenesisState {
	return &GenesisState{
		Bundles: []Bundle{},
	}
}

// Validate performs basic genesis state validation
func (gs GenesisState) Validate() error {
	return nil
}
