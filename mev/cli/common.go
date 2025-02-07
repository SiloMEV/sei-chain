package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func getMEVRpcClient(cmd *cobra.Command) (mev.MevRpcServiceClient, client.Context, error) {

	serverCtx := server.GetServerContextFromCmd(cmd)

	if err := serverCtx.Viper.BindPFlags(cmd.Flags()); err != nil {
		return nil, client.Context{}, err
	}

	mevRpcAddr := serverCtx.Viper.GetString(mev.FlagMEVRpcAddr)

	clientCtx, err := client.GetClientQueryContext(cmd)
	if err != nil {
		return nil, client.Context{}, err
	}

	grpcConn, err := grpc.DialContext(context.Background(), mevRpcAddr, grpc.WithInsecure())
	if err != nil {
		return nil, client.Context{}, err
	}

	mevRpcClient := mev.NewMevRpcServiceClient(grpcConn)

	return mevRpcClient, clientCtx, nil
}
