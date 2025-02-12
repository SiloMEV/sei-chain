package mev_test

import (
	types "github.com/m4ksio/silo-mev-protobuf-go/mev/v1"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/app"
)

func setupKeeper(t *testing.T) (*app.App, tmproto.Header) {
	app := app.Setup(false, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})
	return app, ctx.BlockHeader()
}

func TestKeeper_SubmitAndQueryBundles(t *testing.T) {
	app, _ := setupKeeper(t)

	height := int64(100)

	// Submit a bundle
	bundle := &types.Bundle{
		Transactions: [][]byte{[]byte("tx1"), []byte("tx2")},
		BlockHeight:  uint64(height),
	}

	res, err := app.MevKeeper.AddBundle(height, bundle)
	require.NoError(t, err)
	require.True(t, res)

	// Query bundles
	pending := app.MevKeeper.PendingBundles(height)
	require.Len(t, pending, 1)
	require.Equal(t, bundle.Transactions[0], pending[0].Transactions[0])
}

func TestKeeper_BundlesIdentity(t *testing.T) {
	app, _ := setupKeeper(t)

	height := int64(100)

	// Submit a bundle
	bundle := &types.Bundle{
		Transactions: [][]byte{[]byte("tx1"), []byte("tx2")},
		BlockHeight:  uint64(height),
	}

	res, err := app.MevKeeper.AddBundle(height, bundle)
	require.NoError(t, err)
	require.True(t, res)

	// Query bundles
	pending := app.MevKeeper.PendingBundles(height)
	require.Len(t, pending, 1)
	require.Equal(t, bundle.Transactions[0], pending[0].Transactions[0])

	// sending bundle again won't add it
	res, err = app.MevKeeper.AddBundle(height, bundle)
	require.NoError(t, err)
	require.True(t, res)

	pending = app.MevKeeper.PendingBundles(height)
	require.NoError(t, err)
	require.Equal(t, 1, len(pending))

	//but once we change bundle, it will be stored
	newBundleTx := []byte("tx3")
	bundle.Transactions = [][]byte{newBundleTx}
	res, err = app.MevKeeper.AddBundle(height, bundle)
	require.NoError(t, err)
	require.True(t, res)

	pending = app.MevKeeper.PendingBundles(height)
	require.NoError(t, err)
	require.Equal(t, 2, len(pending))

	require.Contains(t, append(pending[0].Transactions, pending[1].Transactions...), newBundleTx)
}
