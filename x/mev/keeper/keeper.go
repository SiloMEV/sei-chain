package keeper

import (
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

type Keeper struct {
	cdc         codec.BinaryCodec
	ephemeralMu sync.Mutex
	ephemeral   map[string]*types.Bundle
}

func NewKeeper(
	cdc codec.BinaryCodec,
	_ sdk.StoreKey, // keep parameter to maintain compatibility but don't use it
) Keeper {
	return Keeper{
		cdc:       cdc,
		ephemeral: make(map[string]*types.Bundle),
	}
}

// SubmitBundle handles a MsgSubmitBundle
func (k *Keeper) SubmitBundle(ctx sdk.Context, msg *types.MsgSubmitBundle) (*types.MsgSubmitBundleResponse, error) {
	bundle := types.Bundle{
		Sender:    msg.Sender,
		Txs:       msg.Txs,
		BlockNum:  msg.BlockNum,
		Timestamp: msg.Timestamp,
	}

	k.ephemeralMu.Lock()
	k.ephemeral[bundle.Sender] = &bundle
	k.ephemeralMu.Unlock()

	fmt.Println("Submitted bundle for", bundle.Sender, "with", len(bundle.Txs), "transactions")

	return &types.MsgSubmitBundleResponse{
		Success: true,
	}, nil
}
