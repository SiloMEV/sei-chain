package keeper

import (
	"context"
	"fmt"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var _ types.QueryServer = Keeper{}

// PendingBundles implements the Query/PendingBundles gRPC method
func (k Keeper) PendingBundles(c context.Context, req *types.QueryPendingBundlesRequest) (*types.QueryPendingBundlesResponse, error) {
	ctx := sdk.UnwrapSDKContext(c)
	fmt.Println("Querying pending bundles for height", ctx.BlockHeight())
	fmt.Println(ctx.BlockHeader())
	fmt.Println(k.storeKey)
	store := ctx.KVStore(k.storeKey)
	fmt.Println(store)

	var bundles []types.Bundle
	iterator := store.Iterator(nil, nil)
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bundle types.Bundle
		k.cdc.MustUnmarshal(iterator.Value(), &bundle)
		bundles = append(bundles, bundle)
	}

	fmt.Println("Found", len(bundles), "pending bundles from inside the mevkeeper")

	return &types.QueryPendingBundlesResponse{
		Bundles: bundles,
	}, nil
}
