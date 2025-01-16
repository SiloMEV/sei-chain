package ante

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/app/antedecorators"
	"github.com/sei-protocol/sei-chain/x/mev/keeper"
)

type MEVDecorator struct {
	mevKeeper keeper.Keeper
}

func NewMEVDecorator(mk keeper.Keeper) MEVDecorator {
	return MEVDecorator{
		mevKeeper: mk,
	}
}

func (md MEVDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	currentHeight := uint64(ctx.BlockHeight())
	bundle := md.mevKeeper.GetNextBundle(ctx, currentHeight)

	if bundle != nil {
		// Check if this tx is in the bundle
		currentTxBytes := ctx.TxBytes()
		for i, bundledTx := range bundle.Txs {
			if bundledTx == string(currentTxBytes) {
				// Set descending priorities for bundle txs to maintain order
				bundlePriority := antedecorators.MEVBundlePriority - int64(i)
				return ctx.WithPriority(bundlePriority), next(ctx, tx, simulate)
			}
		}
	}

	return next(ctx, tx, simulate)
}
