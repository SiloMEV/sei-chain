package keeper

import (
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

// SubmitBundle submits a new bundle to be included in a future block
func (k Keeper) SubmitBundle(ctx sdk.Context, bundle types.Bundle) error {
	// Basic validation
	if bundle.BlockHeight <= ctx.BlockHeight() {
		return types.ErrInvalidBlockHeight
	}

	if len(bundle.Transactions) > 100 { // Use constant from params
		return types.ErrBundleTooLarge
	}

	// Store the bundle
	k.StoreBundleForBlock(ctx, bundle)

	return nil
}

func (k Keeper) validateBundle(ctx sdk.Context, bundle types.Bundle) error {
	// Check bundle size
	if len(bundle.Transactions) > types.MaxBundleSize {
		return fmt.Errorf("bundle exceeds maximum size")
	}

	// Verify bundle fee meets minimum
	if bundle.BundleFee.Amount.LT(sdk.NewInt(types.MinBundleFee)) {
		return fmt.Errorf("bundle fee too low")
	}

	// Verify target block height
	if bundle.BlockHeight <= ctx.BlockHeight() {
		return fmt.Errorf("invalid target block height")
	}

	return nil
}

func (k Keeper) calculateBundlePriority(ctx sdk.Context, bundle types.Bundle) uint64 {
	// Base priority from fee
	priority := bundle.BundleFee.Amount.Uint64()

	// Additional factors could include:
	// - Searcher reputation
	// - Historical success rate
	// - Bundle size
	// - Time sensitivity

	return priority
}
