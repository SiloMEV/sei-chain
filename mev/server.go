package mev

import (
	"github.com/cosmos/cosmos-sdk/client"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"google.golang.org/grpc"
)

type mevServer struct {
	keeper      *Keeper
	ctxProvider func(int64) sdk.Context
	txConfig    client.TxConfig
	homeDir     string
	server      *grpc.Server
}

//
//func (m mevServer) PendingBundles(_ context.Context, _ *QueryPendingBundlesRequest) (*QueryPendingBundlesResponse, error) {
//	bundles, err := m.keeper.PendingBundles()
//	if err != nil {
//		return nil, err
//	}
//	return &QueryPendingBundlesResponse{Bundles: bundles}, nil
//}
//
//func (m mevServer) SubmitBundle(ctx context.Context, bundle *MsgSubmitBundle) (*MsgSubmitBundleResponse, error) {
//	success, err := m.keeper.SubmitBundle(NewMsgSubmitBundle(bundle))
//	if err != nil {
//		return nil, err
//	}
//	return &MsgSubmitBundleResponse{Success: success}, nil
//}
//
//func StartServer(logger log.Logger,
//	config Config,
//	k *Keeper,
//	ctxProvider func(int64) sdk.Context,
//) error {
//
//	logger.Info("Starting MEV gRPC Server on", "address", config.ListenAddr)
//
//	lis, err := net.Listen("tcp", config.ListenAddr)
//	if err != nil {
//		return err
//	}
//
//	s := grpc.NewServer()
//
//	srv := &mevServer{
//		keeper:      k,
//		ctxProvider: ctxProvider,
//	}
//	RegisterMevRpcServiceServer(s, srv)
//
//	reflection.Register(s)
//
//	err = s.Serve(lis)
//	if err != nil {
//		return err
//	}
//
//	return nil
//}
