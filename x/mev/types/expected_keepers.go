package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// BundleKeeper defines the expected interface for the MEV module keeper
type BundleKeeper interface {
	SubmitBundle(ctx sdk.Context, bundle Bundle) error
	GetBundlesForBlock(ctx sdk.Context, height int64) []Bundle
	StoreBundleForBlock(ctx sdk.Context, bundle Bundle)
	SimulateTx(ctx sdk.Context, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error)
	GetBundleIDForTx(ctx sdk.Context, txHash string) string
	GetBundle(ctx sdk.Context, bundleID string) (Bundle, bool)
}
