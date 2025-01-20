package keeper

import (
	"fmt"

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
	fmt.Println("Submitted bundle for", bundle.Sender, "with", len(bundle.Txs), "transactions")

	return &types.MsgSubmitBundleResponse{
		Success: true,
	}, nil
}
