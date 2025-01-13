package keeper

import (
	"context"

	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var _ types.QueryServer = Keeper{}

// PendingBundles implements the Query/PendingBundles gRPC method
func (k Keeper) PendingBundles(c context.Context, req *types.QueryPendingBundlesRequest) (*types.QueryPendingBundlesResponse, error) {
	// We'll use ctx later when we implement actual bundle storage
	// ctx := sdk.UnwrapSDKContext(c)

	// For now, just return empty bundles
	return &types.QueryPendingBundlesResponse{
		Bundles: []types.Bundle{},
	}, nil
}
