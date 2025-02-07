package mev_test

import (
	"github.com/sei-protocol/sei-chain/mev"
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
	app, header := setupKeeper(t)
	ctx := app.BaseApp.NewContext(false, header)

	// Submit a bundle
	bundle := mev.Bundle{
		Sender:    "test_sender",
		Txs:       [][]byte{[]byte("tx1"), []byte("tx2")},
		BlockNum:  100,
		Timestamp: ctx.BlockTime().Unix(),
	}

	res, err := app.MevKeeper.SubmitBundle(bundle)
	require.NoError(t, err)
	require.True(t, res)

	// Query bundles
	queryRes, err := app.MevKeeper.PendingBundles()
	require.NoError(t, err)
	require.Equal(t, 1, len(queryRes))
	require.Equal(t, bundle.Txs, queryRes[0].Txs)
}
