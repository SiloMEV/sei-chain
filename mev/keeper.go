package mev

import (
	"context"
	"fmt"
	"sync"

	"github.com/cosmos/cosmos-sdk/codec"
	sdk "github.com/cosmos/cosmos-sdk/types"
)

// TODO - separate keeper logic and grpc wrapper? Should we?

type Keeper struct {
	cdc         codec.BinaryCodec
	ephemeralMu sync.Mutex
	ephemeral   map[string]*Bundle
}

func NewKeeper(
	cdc codec.BinaryCodec,
	_ sdk.StoreKey, // keep parameter to maintain compatibility but don't use it
) Keeper {
	return Keeper{
		cdc:       cdc,
		ephemeral: make(map[string]*Bundle),
	}
}

// SubmitBundle handles a MsgSubmitBundle
func (k *Keeper) SubmitBundle(ctx context.Context, msg *MsgSubmitBundle) (*MsgSubmitBundleResponse, error) {
	k.ephemeralMu.Lock()
	defer k.ephemeralMu.Unlock()

	bundle := Bundle{
		Sender:    msg.Sender,
		Txs:       msg.Txs,
		BlockNum:  msg.BlockNum,
		Timestamp: msg.Timestamp,
	}

	k.ephemeral[bundle.Sender] = &bundle

	fmt.Println("Submitted bundle for", bundle.Sender, "with", len(bundle.Txs), "transactions")

	return &MsgSubmitBundleResponse{
		Success: true,
	}, nil
}
