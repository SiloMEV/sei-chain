package mev

import (
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Keeper struct {
	cdc         codec.BinaryCodec
	ephemeralMu sync.Mutex
	ephemeral   map[string]*Bundle
}

func NewKeeper(
	cdc codec.BinaryCodec,
	_ sdk.StoreKey, // keep parameter to maintain compatibility but don't use it
) *Keeper {
	return &Keeper{
		cdc:       cdc,
		ephemeral: make(map[string]*Bundle),
	}
}

// SubmitBundle handles a MsgSubmitBundle
func (k *Keeper) SubmitBundle(bundle Bundle) (bool, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	k.ephemeral[bundle.Sender] = &bundle

	return true, nil
}
