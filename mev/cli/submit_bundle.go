package cli

//func CmdSubmitBundles() *cobra.Command {
//	cmd := &cobra.Command{
//		Use:   "submit-bundle",
//		Short: "Submit new MEV bundle",
//		Args:  cobra.ExactArgs(2),
//		RunE: func(cmd *cobra.Command, args []string) error {
//
//			var txs [][]byte
//			if err := json.Unmarshal([]byte(args[0]), &txs); err != nil {
//				return fmt.Errorf("failed to parse transactions: %w", err)
//			}
//
//			blockNum, ok := math.ParseUint64(args[1])
//			if !ok {
//				return fmt.Errorf("failed to parse block number: %s", args[1])
//			}
//			msgSubmitBundle := mev.MsgSubmitBundle{
//				Sender:    "",
//				Txs:       txs,
//				BlockNum:  blockNum,
//				Timestamp: 0,
//			}
//
//			mevRpcClient, clientCtx, err := getMEVRpcClient(cmd)
//
//			res, err := mevRpcClient.SubmitBundle(context.Background(), &msgSubmitBundle)
//
//			if err != nil {
//				return err
//			}
//
//			return clientCtx.PrintProto(res)
//		},
//	}
//
//	return cmd
//}
