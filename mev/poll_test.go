package mev

import (
	"context"
	"github.com/SiloMEV/silo-mev-protobuf-go/mev/v1"
	"github.com/stretchr/testify/assert"
	"github.com/tendermint/tendermint/libs/log"
	"testing"
)

func TestValidateBundles(t *testing.T) {
	// Mock dependencies
	logger := log.NewNopLogger()
	ctx := context.Background()

	poller := &Poller{
		ctx:           ctx,
		logger:        logger,
		maxBundleSize: 100, // Set a max bundle size for testing
	}

	// Test cases
	tests := []struct {
		name           string
		height         uint64
		bundles        *types.Bundles
		expectedResult []*types.Bundle
	}{
		{
			name:   "valid bundle",
			height: 10,
			bundles: &types.Bundles{
				Bundles: []*types.Bundle{
					{
						BlockHeight: 10,
						Transactions: [][]byte{
							[]byte("transaction1"),
							[]byte("transaction2"),
						},
					},
				},
			},
			expectedResult: []*types.Bundle{
				{
					BlockHeight: 10,
					Transactions: [][]byte{
						[]byte("transaction1"),
						[]byte("transaction2"),
					},
				},
			},
		},
		{
			name:   "bundle with mismatched block height",
			height: 10,
			bundles: &types.Bundles{
				Bundles: []*types.Bundle{
					{
						BlockHeight: 11,
						Transactions: [][]byte{
							[]byte("transaction1"),
						},
					},
				},
			},
			expectedResult: []*types.Bundle{},
		},
		{
			name:   "bundle exceeding max size",
			height: 10,
			bundles: &types.Bundles{
				Bundles: []*types.Bundle{
					{
						BlockHeight: 10,
						Transactions: [][]byte{
							make([]byte, 80), // 80 bytes
							make([]byte, 30), // 30 bytes
						},
					},
				},
			},
			expectedResult: []*types.Bundle{},
		},
		{
			name:   "multiple bundles with mixed validity",
			height: 10,
			bundles: &types.Bundles{
				Bundles: []*types.Bundle{
					{
						BlockHeight: 10,
						Transactions: [][]byte{
							[]byte("tx1"),
						},
					},
					{
						BlockHeight: 11,
						Transactions: [][]byte{
							[]byte("invalidHeight"),
						},
					},
					{
						BlockHeight: 10,
						Transactions: [][]byte{
							make([]byte, 90), // 90 bytes
							make([]byte, 15), // 15 bytes
						},
					},
				},
			},
			expectedResult: []*types.Bundle{
				{
					BlockHeight: 10,
					Transactions: [][]byte{
						[]byte("tx1"),
					},
				},
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			result := poller.validateBundles(tc.height, tc.bundles)
			assert.Equal(t, tc.expectedResult, result)
		})
	}
}
