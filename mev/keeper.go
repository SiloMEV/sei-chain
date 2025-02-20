package mev

import (
	types "github.com/m4ksio/silo-mev-protobuf-go/mev/v1"
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc         codec.BinaryCodec
	ephemeralMu sync.Mutex

	// TODO rethink pointers
	// block heigh -> bundleID -> bundle
	ephemeral map[int64]map[string]*types.Bundle
	minHeight int64
}

func NewKeeper(
	cdc codec.BinaryCodec,
	_ sdk.StoreKey, // keep parameter to maintain compatibility but don't use it
) *Keeper {
	return &Keeper{
		cdc:       cdc,
		ephemeral: make(map[int64]map[string]*types.Bundle),
		minHeight: -1,
	}
}

func (k *Keeper) AddBundle(height int64, bundle *types.Bundle) (bool, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	if k.ephemeral[int64(height)] == nil {
		k.ephemeral[int64(height)] = make(map[string]*types.Bundle)
	}

	k.ephemeral[int64(height)][bundle.ID()] = bundle

	if k.minHeight == -1 {
		k.minHeight = int64(height)
	}

	return true, nil
}

// TODO keep error as it probably will be useful if we keep track of expired bundles etc
func (k *Keeper) AddBundles(height int64, bundles []*types.Bundle) (bool, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	if k.ephemeral[int64(height)] == nil {
		k.ephemeral[int64(height)] = make(map[string]*types.Bundle)
	}

	for _, bundle := range bundles {
		k.ephemeral[int64(height)][bundle.ID()] = bundle
	}

	if k.minHeight == -1 {
		k.minHeight = int64(height)
	}

	return true, nil
}

func (k *Keeper) PendingBundles(height int64) []*types.Bundle {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	var bundles []*types.Bundle
	for _, b := range k.ephemeral[height] {

		bundles = append(bundles, b)

	}

	return bundles
}

// TODO rethink this logic
func (k *Keeper) DropBundlesBelow(height int64) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	if k.minHeight == -1 {
		return
	}

	for i := height; i < k.minHeight; i-- {
		delete(k.ephemeral, i)
	}

	k.minHeight = height
}
