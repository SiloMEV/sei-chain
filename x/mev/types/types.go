package types

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
)

type Bundle struct {
	Transactions [][]byte       // Raw transactions in the bundle
	BlockHeight  int64          // Target block height
	Sender       sdk.AccAddress // Address of the searcher
	BundleFee    sdk.Coin       // Fee paid for bundle inclusion
	Priority     uint64         // Bundle priority score
}

type BundleStatus struct {
	Included    bool   `json:"included"`
	BlockHeight int64  `json:"block_height"`
	Error       string `json:"error"`
}

const (
	MaxBundleSize      = 50   // Maximum number of transactions in a bundle
	MaxBundlesPerBlock = 10   // Maximum bundles per block
	MinBundleFee       = 1000 // Minimum fee required for bundle submission
)

// Bundles is a collection of Bundle objects for codec marshaling
type Bundles struct {
	Items []Bundle `json:"items"`
}
