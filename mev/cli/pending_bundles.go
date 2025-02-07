package cli

import (
	"context"
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/spf13/cobra"
)

func CmdPendingBundles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-bundles",
		Short: "Query pending MEV bundles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			mevRpcClient, clientCtx, err := getMEVRpcClient(cmd)

			res, err := mevRpcClient.PendingBundles(context.Background(), &mev.QueryPendingBundlesRequest{})
			if err != nil {
				return err
			}

			return clientCtx.PrintProto(&mev.QueryPendingBundlesResponse{Bundles: res.Bundles})
		},
	}

	return cmd
}
