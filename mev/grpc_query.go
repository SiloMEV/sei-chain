package mev

import (
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

var _ types.QueryServer = Keeper{}

func (k Keeper) PendingBundles() ([]Bundle, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	bundles := make([]Bundle, 0, len(k.ephemeral))
	for _, bundle := range k.ephemeral {
		bundles = append(bundles, *bundle)
	}

	return bundles, nil
}
