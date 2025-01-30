package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var _ types.QueryServer = Keeper{}

// PendingBundles implements the Query/PendingBundles gRPC method
func (k Keeper) PendingBundles(c context.Context, req *types.QueryPendingBundlesRequest) (*types.QueryPendingBundlesResponse, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	bundles := make([]types.Bundle, 0, len(k.ephemeral))
	for _, bundle := range k.ephemeral {
		bundles = append(bundles, *bundle)
	}

	return &types.QueryPendingBundlesResponse{
		Bundles: bundles,
	}, nil
}
