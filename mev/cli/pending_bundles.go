package cli

import (
	"fmt"
	"github.com/spf13/cobra"
)

func CmdPendingBundles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "pending-bundles",
		Short: "Query pending MEV bundles",
		Args:  cobra.NoArgs,
		RunE: func(cmd *cobra.Command, args []string) error {

			fmt.Println("Not implemented yet")

			//mevRpcClient, clientCtx, err := getMEVRpcClient(cmd)
			//
			//res, err := mevRpcClient.PendingBundles(context.Background(), &mev.QueryPendingBundlesRequest{})
			//if err != nil {
			//	return err
			//}
			//
			//return clientCtx.PrintProto(&mev.QueryPendingBundlesResponse{Bundles: res.Bundles})

			return nil
		},
	}

	return cmd
}
