package mev

import (
	"encoding/json"
	"fmt"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/codec"
	cdctypes "github.com/cosmos/cosmos-sdk/codec/types"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	"github.com/gorilla/mux"
	"github.com/grpc-ecosystem/grpc-gateway/runtime"
	"github.com/spf13/cobra"
	abci "github.com/tendermint/tendermint/abci/types"

	"github.com/sei-protocol/sei-chain/x/mev/keeper"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var (
	_ module.AppModule      = AppModule{}
	_ module.AppModuleBasic = AppModuleBasic{}
)

// AppModuleBasic defines the basic application module used by the MEV module.
type AppModuleBasic struct{}

func (AppModuleBasic) Name() string {
	return types.ModuleName
}

func (AppModuleBasic) RegisterLegacyAminoCodec(cdc *codec.LegacyAmino) {
	types.RegisterCodec(cdc)
}

func (AppModuleBasic) RegisterInterfaces(registry cdctypes.InterfaceRegistry) {
	types.RegisterInterfaces(registry)
}

func (AppModuleBasic) DefaultGenesis(cdc codec.JSONCodec) json.RawMessage {
	defaultState := make(map[string]json.RawMessage)
	bz, err := json.Marshal(defaultState)
	if err != nil {
		panic(err)
	}
	return bz
}

func (AppModuleBasic) ValidateGenesis(cdc codec.JSONCodec, _ client.TxEncodingConfig, bz json.RawMessage) error {
	var data map[string]json.RawMessage
	err := json.Unmarshal(bz, &data)
	if err != nil {
		return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
	}
	return nil
}

func (AppModuleBasic) RegisterRESTRoutes(_ client.Context, _ *mux.Router) {}

func (AppModuleBasic) RegisterGRPCGatewayRoutes(_ client.Context, _ *runtime.ServeMux) {}

func (AppModuleBasic) GetTxCmd() *cobra.Command {
	return nil // Implement if needed
}

func (AppModuleBasic) GetQueryCmd() *cobra.Command {
	return nil // Implement if needed
}

func (AppModuleBasic) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, stream <-chan json.RawMessage) error {
	for bz := range stream {
		var data map[string]json.RawMessage
		if err := json.Unmarshal(bz, &data); err != nil {
			return fmt.Errorf("failed to unmarshal %s genesis state: %w", types.ModuleName, err)
		}
	}
	return nil
}

// AppModule implements an application module for the MEV module.
type AppModule struct {
	AppModuleBasic
	keeper keeper.Keeper
}

func NewAppModule(k keeper.Keeper) AppModule {
	return AppModule{
		AppModuleBasic: AppModuleBasic{},
		keeper:         k,
	}
}

func (AppModule) Name() string {
	return types.ModuleName
}

func (am AppModule) RegisterInvariants(_ sdk.InvariantRegistry) {}

func (am AppModule) Route() sdk.Route {
	return sdk.NewRoute(types.RouterKey, nil) // Implement handler if needed
}

func (am AppModule) QuerierRoute() string {
	return types.ModuleName
}

func (am AppModule) LegacyQuerierHandler(_ *codec.LegacyAmino) sdk.Querier {
	return nil
}

func (am AppModule) RegisterServices(_ module.Configurator) {}

func (am AppModule) InitGenesis(ctx sdk.Context, cdc codec.JSONCodec, data json.RawMessage) []abci.ValidatorUpdate {
	var genesisState map[string]json.RawMessage
	json.Unmarshal(data, &genesisState)
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ExportGenesis(ctx sdk.Context, cdc codec.JSONCodec) json.RawMessage {
	genesisState := make(map[string]json.RawMessage)
	bz, err := json.Marshal(genesisState)
	if err != nil {
		panic(err)
	}
	return bz
}

func (am AppModule) ExportGenesisStream(ctx sdk.Context, cdc codec.JSONCodec) <-chan json.RawMessage {
	stream := make(chan json.RawMessage)
	go func() {
		defer close(stream)
		genesisState := make(map[string]json.RawMessage)
		bz, err := json.Marshal(genesisState)
		if err != nil {
			panic(err)
		}
		stream <- bz
	}()
	return stream
}

func (am AppModule) ConsensusVersion() uint64 {
	return 1
}

func (am AppModule) BeginBlock(_ sdk.Context, _ abci.RequestBeginBlock) {}

func (am AppModule) EndBlock(_ sdk.Context, _ abci.RequestEndBlock) []abci.ValidatorUpdate {
	return []abci.ValidatorUpdate{}
}

func (am AppModule) ValidateGenesisStream(cdc codec.JSONCodec, config client.TxEncodingConfig, stream <-chan json.RawMessage) error {
	return am.AppModuleBasic.ValidateGenesisStream(cdc, config, stream)
}
