package mev

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc/credentials"

	types "github.com/SiloMEV/silo-mev-protobuf-go/mev/v1"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"
)

type Poller struct {
	client            types.BundleProviderClient
	keeper            *Keeper
	lastBlockProvider func() int64
	logger            log.Logger
	ctx               context.Context
	maxBundleSize     uint64
}

func (p *Poller) run() {

	lastHeight := p.lastBlockProvider()

	bundles, err := p.client.GetBundles(context.Background(), &types.GetBundlesRequest{MinBlockHeight: uint64(lastHeight)})
	if err != nil {
		p.logger.Error("Error while querying bundle server", "err", err)
		return
	}
	for height, bundles := range bundles.Bundles {

		validBundles := p.validateBundles(height, bundles)

		p.keeper.SetBundles(int64(height), validBundles)
	}

	p.keeper.DropBundlesAtAndBelow(lastHeight - 1)
}

func (p *Poller) validateBundles(height uint64, bundles *types.Bundles) []*types.Bundle {

	validBundles := make([]*types.Bundle, 0, len(bundles.Bundles))

	for _, bundle := range bundles.Bundles {
		if bundle.BlockHeight != height {
			p.logger.Debug("Bundle block height does not match height of container", "bundleHeight", bundle.BlockHeight, "containerHeight", height)
			continue
		}
		var bundleSize uint64 = 0
		for _, tx := range bundle.Transactions {
			bundleSize += uint64(len(tx))
		}
		if bundleSize > p.maxBundleSize {
			p.logger.Debug("Bundle size exceeds max size", "bundleSize", bundleSize, "maxBundleSize", p.maxBundleSize)
			continue
		}

		validBundles = append(validBundles, bundle)
	}

	return validBundles
}

func NewPoller(ctx context.Context, logger log.Logger, config Config, keeper *Keeper, lastBlockProvider func() int64, maxBundleSize int64) (*Poller, error) {

	logger.Info("Starting bundle provider poller")

	if config.CertFile == "" && !config.Insecure {
		return nil, fmt.Errorf("either certFile or insecure must be set")
	}

	var option grpc.DialOption

	if config.CertFile != "" {
		creds, err := credentials.NewClientTLSFromFile(config.CertFile, "")
		if err != nil {
			return nil, fmt.Errorf("error while loading TLS certificate: %w", err)
		}
		option = grpc.WithTransportCredentials(creds)
	} else {
		option = grpc.WithInsecure()
	}

	grpcConn, err := grpc.DialContext(ctx, config.ServerAddr, option)
	if err != nil {
		return nil, err
	}

	bundleProviderClient := types.NewBundleProviderClient(grpcConn)

	p := &Poller{
		client:            bundleProviderClient,
		keeper:            keeper,
		lastBlockProvider: lastBlockProvider,
		logger:            logger,
		ctx:               ctx,
		maxBundleSize:     uint64(maxBundleSize),
	}

	ticker := time.NewTicker(config.PollInterval)

	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case <-ticker.C:
				p.run()
			}
		}
	}()

	return p, nil
}
