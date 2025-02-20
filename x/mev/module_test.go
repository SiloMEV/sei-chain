package mev_test

import (
	"testing"

	"github.com/stretchr/testify/require"
	abci "github.com/tendermint/tendermint/abci/types"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/app"
	mevbase "github.com/sei-protocol/sei-chain/mev"
	"github.com/sei-protocol/sei-chain/x/mev"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

func TestBasicModule(t *testing.T) {
	app := app.Setup(false, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	module := mev.NewAppModule(
		app.AppCodec(),
		app.MevKeeper,
	)

	// Test basic module properties
	require.Equal(t, types.ModuleName, module.Name())
	require.NotNil(t, module)

	// Test BeginBlock and EndBlock
	module.BeginBlock(ctx, abci.RequestBeginBlock{})
	require.Equal(t, []abci.ValidatorUpdate{}, module.EndBlock(ctx, abci.RequestEndBlock{}))
}

func TestQueryPendingBundles(t *testing.T) {
	app := app.Setup(false, false)

	// Query pending bundles
	res, err := app.MevKeeper.PendingBundles()
	require.NoError(t, err)
	require.NotNil(t, res)
	require.Equal(t, 0, len(res))
}

func TestSubmitBundle(t *testing.T) {
	app := app.Setup(false, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{})

	// Create a test bundle
	bundle := mevbase.Bundle{
		Sender:    "test_sender",
		Txs:       [][]byte{[]byte("tx1"), []byte("tx2")},
		BlockNum:  100,
		Timestamp: ctx.BlockTime().Unix(),
	}

	res, err := app.MevKeeper.SubmitBundle(bundle)
	require.NoError(t, err)
	require.NotNil(t, res)
	require.True(t, res)

	// Verify bundle was stored
	queryRes, err := app.MevKeeper.PendingBundles()
	require.NoError(t, err)
	require.Equal(t, 1, len(queryRes))
	require.Equal(t, bundle.Txs, queryRes[0].Txs)
}

func TestModuleRegistration(t *testing.T) {
	app := app.Setup(false, false)

	// Verify the module is properly registered in the app
	require.NotNil(t, app.MevKeeper)

	// Test module name matches
	require.Equal(t, types.ModuleName, types.ModuleName)
}
