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
	// k.SetBundle(ctx, &msg.Bundle)

	return &types.MsgSubmitBundleResponse{
		Success: true,
	}, nil
}
