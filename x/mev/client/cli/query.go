package cli

// GetQueryCmd returns the cli query commands for the mev module
//func GetQueryCmd() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:                        types.ModuleName,
//		Short:                      fmt.Sprintf("Querying commands for the %s module", types.ModuleName),
//		DisableFlagParsing:         true,
//		SuggestionsMinimumDistance: 2,
//		RunE:                       client.ValidateCmd,
//	}
//
//	cmd.AddCommand(
//		CmdQueryPendingBundles(),
//	)
//
//	return cmd
//}

//func CmdQueryPendingBundles() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "pending-bundles",
//		Short: "Query pending MEV bundles",
//		Args:  cobra.NoArgs,
//		RunE: func(cmd *cobra.Command, args []string) error {
//			clientCtx, err := client.GetClientQueryContext(cmd)
//			if err != nil {
//				return err
//			}
//
//			// Get gRPC connection from client context
//			queryClient := types.NewQueryClient(clientCtx)
//
//			res, err := queryClient.PendingBundles(context.Background(), &types.QueryPendingBundlesRequest{})
//			if err != nil {
//				return err
//			}
//
//			return clientCtx.PrintProto(&types.QueryPendingBundlesResponse{Bundles: res.Bundles})
//		},
//	}
//
//	flags.AddQueryFlagsToCmd(cmd)
//	return cmd
//}
