package keeper

import (
	"context"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var _ types.QueryServer = Keeper{}

// PendingBundles implements the Query/PendingBundles gRPC method
func (k Keeper) PendingBundles(c context.Context, req *types.QueryPendingBundlesRequest) (*types.QueryPendingBundlesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	store := ctx.KVStore(k.storeKey)

	var bundles []types.Bundle
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bundle types.Bundle
		k.cdc.MustUnmarshal(iterator.Value(), &bundle)
		bundles = append(bundles, bundle)
	}

	return &types.QueryPendingBundlesResponse{
		Bundles: bundles,
	}, nil
}
