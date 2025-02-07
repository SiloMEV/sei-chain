package cli

import (
	"context"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func CmdPendingBundles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-bundles",
		Short: "Query pending MEV bundles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			serverCtx := server.GetServerContextFromCmd(cmd)

			if err := serverCtx.Viper.BindPFlags(cmd.Flags()); err != nil {
				return err
			}

			mevRpcAddr := serverCtx.Viper.GetString(mev.FlagMEVRpcAddr)

			clientCtx, err := client.GetClientQueryContext(cmd)
			if err != nil {
				return err
			}

			grpcConn, err := grpc.DialContext(context.Background(), mevRpcAddr, grpc.WithInsecure())
			if err != nil {
				return err
			}

			// Get gRPC connection from client context
			mevRpcClient := mev.NewMevRpcServiceClient(grpcConn)

			res, err := mevRpcClient.PendingBundles(context.Background(), &mev.QueryPendingBundlesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&mev.QueryPendingBundlesResponse{Bundles: res.Bundles})
		},
	}

	return cmd
}
