package keeper_test

import (
	"testing"
	"time"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/stretchr/testify/require"

	"github.com/sei-protocol/sei-chain/testutil"
	"github.com/sei-protocol/sei-chain/x/mev/keeper"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

func TestBundleSubmission(t *testing.T) {
	k, ctx := setupKeeper(t)
	addr := testutil.GenerateAddress()

	t.Run("successful bundle submission", func(t *testing.T) {
		bundle := types.Bundle{
			Transactions: [][]byte{[]byte("tx1"), []byte("tx2")},
			BlockHeight:  ctx.BlockHeight() + 1,
			Sender:       addr,
			BundleFee:    sdk.NewCoin("usei", sdk.NewInt(1000)),
			Timestamp:    time.Now().Unix(),
		}

		err := k.SubmitBundle(ctx, bundle)
		require.NoError(t, err)

		// Verify bundle was stored
		storedBundle := k.GetBundlesForBlock(ctx, bundle.BlockHeight)
		require.Len(t, storedBundle, 1)
		require.Equal(t, bundle.Transactions, storedBundle[0].Transactions)
	})

	t.Run("bundle with invalid block height", func(t *testing.T) {
		bundle := types.Bundle{
			Transactions: [][]byte{[]byte("tx1")},
			BlockHeight:  ctx.BlockHeight() - 1, // Past block height
			Sender:       addr,
			BundleFee:    sdk.NewCoin("usei", sdk.NewInt(1000)),
			Timestamp:    time.Now().Unix(),
		}

		err := k.SubmitBundle(ctx, bundle)
		require.Error(t, err)
		require.Contains(t, err.Error(), "invalid target block height")
	})

	t.Run("bundle exceeding max size", func(t *testing.T) {
		// Create bundle with too many transactions
		txs := make([][]byte, types.MaxBundleSize+1)
		for i := range txs {
			txs[i] = []byte("tx")
		}

		bundle := types.Bundle{
			Transactions: txs,
			BlockHeight:  ctx.BlockHeight() + 1,
			Sender:       addr,
			BundleFee:    sdk.NewCoin("usei", sdk.NewInt(1000)),
			Timestamp:    time.Now().Unix(),
		}

		err := k.SubmitBundle(ctx, bundle)
		require.Error(t, err)
		require.Contains(t, err.Error(), "bundle exceeds maximum size")
	})

	t.Run("bundle with insufficient fee", func(t *testing.T) {
		bundle := types.Bundle{
			Transactions: [][]byte{[]byte("tx1")},
			BlockHeight:  ctx.BlockHeight() + 1,
			Sender:       addr,
			BundleFee:    sdk.NewCoin("usei", sdk.NewInt(1)), // Too low
			Timestamp:    time.Now().Unix(),
		}

		err := k.SubmitBundle(ctx, bundle)
		require.Error(t, err)
		require.Contains(t, err.Error(), "bundle fee too low")
	})
}

func TestBundleProcessing(t *testing.T) {
	k, ctx := setupKeeper(t)
	addr := testutil.GenerateAddress()

	t.Run("process bundles in priority order", func(t *testing.T) {
		// Submit multiple bundles with different priorities
		bundles := []types.Bundle{
			{
				Transactions: [][]byte{[]byte("tx1")},
				BlockHeight:  ctx.BlockHeight() + 1,
				Sender:       addr,
				BundleFee:    sdk.NewCoin("usei", sdk.NewInt(1000)),
				Priority:     1,
			},
			{
				Transactions: [][]byte{[]byte("tx2")},
				BlockHeight:  ctx.BlockHeight() + 1,
				Sender:       addr,
				BundleFee:    sdk.NewCoin("usei", sdk.NewInt(2000)),
				Priority:     2,
			},
		}

		for _, bundle := range bundles {
			err := k.SubmitBundle(ctx, bundle)
			require.NoError(t, err)
		}

		// Process bundles
		processedTxs := k.ProcessBundlesForBlock(ctx)

		// Verify order based on priority
		require.Len(t, processedTxs, 2)
		require.Equal(t, []byte("tx2"), processedTxs[0]) // Higher priority should be first
		require.Equal(t, []byte("tx1"), processedTxs[1])
	})
}

func TestPriorityDecorator(t *testing.T) {
	k, ctx := setupKeeper(t)
	decorator := antedecorators.NewPriorityDecorator(k)

	t.Run("mev bundle transaction priority", func(t *testing.T) {
		// Create and submit a bundle
		bundle := types.Bundle{
			Transactions: [][]byte{[]byte("tx1")},
			BlockHeight:  ctx.BlockHeight() + 1,
			Priority:     100,
		}
		err := k.SubmitBundle(ctx, bundle)
		require.NoError(t, err)

		// Create transaction that's part of the bundle
		tx := testutil.NewMockTx([]byte("tx1"))

		// Run through priority decorator
		newCtx, err := decorator.AnteHandle(ctx, tx, false, testutil.NoopAnteHandler())
		require.NoError(t, err)

		// Verify priority was set correctly
		require.Equal(t, int64(antedecorators.MEVBundlePriority+100), newCtx.Priority())
	})

	t.Run("standard transaction priority", func(t *testing.T) {
		tx := testutil.NewMockTx([]byte("regular_tx"))

		newCtx, err := decorator.AnteHandle(ctx, tx, false, testutil.NoopAnteHandler())
		require.NoError(t, err)

		require.Equal(t, int64(antedecorators.StandardPriority), newCtx.Priority())
	})
}

// Helper function to setup keeper and context for testing
func setupKeeper(t *testing.T) (keeper.Keeper, sdk.Context) {
	// Create in-memory database
	db := testutil.NewInMemoryDB()

	// Create test app context
	ctx := testutil.NewContext(db)

	// Initialize and return keeper
	k := keeper.NewKeeper(
		testutil.NewTestCodec(),
		testutil.NewTestStoreKey(types.StoreKey),
		testutil.NewMockBankKeeper(),
	)

	return k, ctx
}
