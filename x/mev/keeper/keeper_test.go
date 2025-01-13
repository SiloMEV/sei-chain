package keeper_test

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
	tmproto "github.com/tendermint/tendermint/proto/tendermint/types"

	"github.com/sei-protocol/sei-chain/app"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

func setupKeeper(t testing.TB) (*app.App, tmproto.Header) {
	app := app.Setup(false, false)
	ctx := app.BaseApp.NewContext(false, tmproto.Header{Time: time.Now()})
	return app, ctx.BlockHeader()
}

func TestKeeper_SubmitAndQueryBundles(t *testing.T) {
	app, header := setupKeeper(t)
	ctx := app.BaseApp.NewContext(false, header)

	// Submit a bundle
	bundle := types.Bundle{
		Sender:    []byte("test_sender"),
		Txs:       []string{"tx1", "tx2"},
		BlockNum:  100,
		Timestamp: ctx.BlockTime().Unix(),
	}

	msg := types.NewMsgSubmitBundle(bundle)
	res, err := app.MevKeeper.SubmitBundle(ctx, msg)
	require.NoError(t, err)
	require.True(t, res.Success)

	// Query bundles
	queryRes, err := app.MevKeeper.PendingBundles(context.Background(), &types.QueryPendingBundlesRequest{})
	require.NoError(t, err)
	require.Equal(t, 1, len(queryRes.Bundles))
	require.Equal(t, bundle.Txs, queryRes.Bundles[0].Txs)
}
