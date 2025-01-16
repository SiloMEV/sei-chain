package keeper

import (
	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

type Keeper struct {
	cdc      codec.BinaryCodec
	storeKey sdk.StoreKey
}

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
) Keeper {
	return Keeper{
		cdc:      cdc,
		storeKey: storeKey,
	}
}

// SubmitBundle handles a MsgSubmitBundle
func (k Keeper) SubmitBundle(ctx sdk.Context, msg *types.MsgSubmitBundle) (*types.MsgSubmitBundleResponse, error) {
	// Store the bundle
	bundle := types.Bundle{
		Sender:    msg.Sender,
		Txs:       msg.Txs,
		BlockNum:  msg.BlockNum,
		Timestamp: msg.Timestamp,
	}

	// TODO: Add actual bundle storage logic here
	store := ctx.KVStore(k.storeKey)
	bz := k.cdc.MustMarshal(&bundle)
	store.Set([]byte(bundle.Sender), bz)

	return &types.MsgSubmitBundleResponse{
		Success: true,
	}, nil
}

func (k Keeper) GetNextBundle(ctx sdk.Context, blockHeight uint64) *types.Bundle {
	store := ctx.KVStore(k.storeKey)

	var nextBundle *types.Bundle
	iterator := sdk.KVStorePrefixIterator(store, []byte("bundle"))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bundle types.Bundle
		k.cdc.MustUnmarshal(iterator.Value(), &bundle)
		if bundle.BlockNum == blockHeight {
			nextBundle = &bundle
			break
		}
	}

	return nextBundle
}

func (k Keeper) CleanupProcessedBundles(ctx sdk.Context) {
	store := ctx.KVStore(k.storeKey)
	currentHeight := uint64(ctx.BlockHeight())

	iterator := sdk.KVStorePrefixIterator(store, []byte("bundle"))
	defer iterator.Close()

	for ; iterator.Valid(); iterator.Next() {
		var bundle types.Bundle
		k.cdc.MustUnmarshal(iterator.Value(), &bundle)
		if bundle.BlockNum <= currentHeight {
			store.Delete(iterator.Key())
		}
	}
}
