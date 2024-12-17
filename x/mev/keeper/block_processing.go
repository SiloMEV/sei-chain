package keeper

import (
	"sort"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

// ProcessBundlesForBlock processes all bundles for the current block
func (k Keeper) ProcessBundlesForBlock(ctx sdk.Context) []sdk.Tx {
	height := ctx.BlockHeight()
	bundles := k.GetBundlesForBlock(ctx, height)

	// Sort bundles by priority
	sort.Slice(bundles, func(i, j int) bool {
		return bundles[i].Priority > bundles[j].Priority
	})

	var processedTxs []sdk.Tx
	for _, bundle := range bundles {
		success, txs := k.processSingleBundle(ctx, bundle)
		if success {
			processedTxs = append(processedTxs, txs...)
		}
	}

	return processedTxs
}

func (k Keeper) processSingleBundle(ctx sdk.Context, bundle types.Bundle) (bool, []sdk.Tx) {
	var txs []sdk.Tx

	// Decode transactions
	for _, txBytes := range bundle.Transactions {
		tx, err := k.txDecoder(txBytes)
		if err != nil {
			return false, nil
		}
		txs = append(txs, tx)
	}

	// Simulate bundle execution
	cacheCtx, _ := ctx.CacheContext()
	for _, tx := range txs {
		_, result, err := k.SimulateTx(cacheCtx, tx)
		if err != nil || result == nil {
			return false, nil
		}
	}

	return true, txs
}
