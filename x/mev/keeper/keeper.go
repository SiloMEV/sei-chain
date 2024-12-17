package keeper

import (
	"encoding/binary"
	"encoding/json"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

// Keeper implements types.BundleKeeper interface
type Keeper struct {
	cdc       codec.BinaryCodec
	storeKey  sdk.StoreKey
	memKey    sdk.StoreKey
	txDecoder sdk.TxDecoder
}

// Ensure Keeper implements BundleKeeper interface
var _ types.BundleKeeper = (*Keeper)(nil)

func NewKeeper(
	cdc codec.BinaryCodec,
	storeKey sdk.StoreKey,
	memKey sdk.StoreKey,
	txDecoder sdk.TxDecoder,
) Keeper {
	return Keeper{
		cdc:       cdc,
		storeKey:  storeKey,
		memKey:    memKey,
		txDecoder: txDecoder,
	}
}

// GetBundlesForBlock returns all bundles for a specific block height
func (k Keeper) GetBundlesForBlock(ctx sdk.Context, height int64) []types.Bundle {
	store := ctx.KVStore(k.storeKey)
	key := getBundleKeyForHeight(height)

	bz := store.Get(key)
	if bz == nil {
		return nil
	}

	var bundles []types.Bundle
	err := json.Unmarshal(bz, &bundles)
	if err != nil {
		panic(err)
	}
	return bundles
}

// StoreBundleForBlock stores a bundle for a specific block
func (k Keeper) StoreBundleForBlock(ctx sdk.Context, bundle types.Bundle) {
	store := ctx.KVStore(k.storeKey)
	key := getBundleKeyForHeight(bundle.BlockHeight)

	bundles := k.GetBundlesForBlock(ctx, bundle.BlockHeight)
	bundles = append(bundles, bundle)

	bz, err := json.Marshal(bundles)
	if err != nil {
		panic(err)
	}
	store.Set(key, bz)
}

// Helper functions for keys
func getBundleKeyForHeight(height int64) []byte {
	heightBz := make([]byte, 8)
	binary.BigEndian.PutUint64(heightBz, uint64(height))
	return append([]byte("bundle:"), heightBz...)
}

// SimulateTx simulates a transaction
func (k Keeper) SimulateTx(ctx sdk.Context, tx sdk.Tx) (sdk.GasInfo, *sdk.Result, error) {
	gasInfo := sdk.GasInfo{GasWanted: 0, GasUsed: 0}
	result := &sdk.Result{}

	// In a real implementation, you'd want to properly handle gas estimation
	// and transaction simulation. This is a simplified version.
	return gasInfo, result, nil
}

// GetBundleIDForTx returns the bundle ID for a given transaction
func (k Keeper) GetBundleIDForTx(ctx sdk.Context, txHash string) string {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte("tx_to_bundle:"), []byte(txHash)...)
	return string(store.Get(key))
}

// GetBundle returns a bundle by its ID
func (k Keeper) GetBundle(ctx sdk.Context, bundleID string) (types.Bundle, bool) {
	store := ctx.KVStore(k.storeKey)
	key := append([]byte("bundle:"), []byte(bundleID)...)

	bz := store.Get(key)
	if bz == nil {
		return types.Bundle{}, false
	}

	var bundle types.Bundle
	err := json.Unmarshal(bz, &bundle)
	if err != nil {
		panic(err)
	}
	return bundle, true
}
