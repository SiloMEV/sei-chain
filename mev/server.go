package mev

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/tendermint/tendermint/libs/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"

	"net"
)

type mevServer struct {
	keeper      *Keeper
	ctxProvider func(int64) sdk.Context
	txConfig    client.TxConfig
	homeDir     string
	server      *grpc.Server
}

func (m mevServer) PendingBundles(ctx context.Context, request *QueryPendingBundlesRequest) (*QueryPendingBundlesResponse, error) {
	return m.keeper.PendingBundles(ctx, request)
}

func (m mevServer) SubmitBundle(ctx context.Context, bundle *MsgSubmitBundle) (*MsgSubmitBundleResponse, error) {
	return m.keeper.SubmitBundle(ctx, bundle)
}

func StartServer(logger log.Logger,
	config Config,
	k *Keeper,
	ctxProvider func(int64) sdk.Context,
) error {

	logger.Info("Starting MEV gRPC Server on", "address", config.ListenAddr)

	lis, err := net.Listen("tcp", config.ListenAddr)
	if err != nil {
		return err
	}

	s := grpc.NewServer()

	srv := &mevServer{
		keeper:      k,
		ctxProvider: ctxProvider,
	}
	RegisterMevRpcServiceServer(s, srv)

	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		return err
	}

	return nil
}
