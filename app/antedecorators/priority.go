package antedecorators

import (
	"crypto/sha256"
	"encoding/hex"
	"math"

	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/keeper"
)

const (
	MaxPriority          = math.MaxInt64
	SystemPriority       = math.MaxInt64 - 1
	MEVBundlePriority    = math.MaxInt64 - 50
	OraclePriority       = math.MaxInt64 - 100
	EVMAssociatePriority = math.MaxInt64 - 500
	StandardPriority     = math.MaxInt64 - 1000
)

type PriorityDecorator struct {
	mevKeeper keeper.Keeper
	txConfig  client.TxConfig
}

func NewPriorityDecorator(mk keeper.Keeper, txConfig client.TxConfig) PriorityDecorator {
	return PriorityDecorator{
		mevKeeper: mk,
		txConfig:  txConfig,
	}
}

func (pd PriorityDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	// Get tx hash
	txBytes, err := pd.txConfig.TxEncoder()(tx)
	if err != nil {
		return ctx, err
	}

	hash := sha256.Sum256(txBytes)
	txHash := hex.EncodeToString(hash[:])

	// Check if tx is part of a bundle
	bundleID := pd.mevKeeper.GetBundleIDForTx(ctx, txHash)
	if bundleID != "" {
		if _, found := pd.mevKeeper.GetBundle(ctx, bundleID); found {
			ctx = ctx.WithPriority(MEVBundlePriority)
			return next(ctx, tx, simulate)
		}
	}

	ctx = ctx.WithPriority(StandardPriority)
	return next(ctx, tx, simulate)
}
