package mev

import (
	"context"
	types "github.com/m4ksio/silo-mev-protobuf-go/mev/v1"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"
	"time"
)

type Poller struct {
	client            types.BundleProviderClient
	keeper            *Keeper
	lastBlockProvider func() int64
	logger            log.Logger
}

func (p *Poller) run() {

	lastHeight := p.lastBlockProvider()

	bundles, err := p.client.GetBundles(context.Background(), &types.GetBundlesRequest{MinBlockHeight: uint64(lastHeight)})
	if err != nil {
		p.logger.Error("Error while querying SILO server for bundles", "err", err)
		return
	}
	// TODO validate data, don't trust height
	for height, bundles := range bundles.Bundles {
		p.keeper.AddBundles(int64(height), bundles.Bundles)
	}

}

func NewPoller(logger log.Logger, config Config, keeper *Keeper, lastBlockProvider func() int64) (*Poller, error) {

	logger.Info("Starting bundle provider poller")

	// TODO secure grpc connection
	grpcConn, err := grpc.DialContext(context.Background(), config.ServerAddr, grpc.WithInsecure())
	if err != nil {
		return nil, err
	}

	bundleProviderClient := types.NewBundleProviderClient(grpcConn)

	p := &Poller{
		client:            bundleProviderClient,
		keeper:            keeper,
		lastBlockProvider: lastBlockProvider,
		logger:            logger,
	}

	ticker := time.NewTicker(200 * time.Millisecond)

	go func() {
		for range ticker.C {
			p.run()
		}
	}()

	return p, nil
}
