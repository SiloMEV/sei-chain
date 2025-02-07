package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/server"
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/spf13/cobra"
	"google.golang.org/grpc"
)

func CmdSubmitBundles() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-bundle",
		Short: "Submit new MEV bundle",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {

			var txs [][]byte
			if err := json.Unmarshal([]byte(args[0]), &txs); err != nil {
				return fmt.Errorf("failed to parse transactions: %w", err)
			}

			blockNum, ok := math.ParseUint64(args[1])
			if !ok {
				return fmt.Errorf("failed to parse block number: %s", args[1])
			}
			serverCtx := server.GetServerContextFromCmd(cmd)

			msgSubmitBundle := mev.MsgSubmitBundle{
				Sender:    "",
				Txs:       txs,
				BlockNum:  blockNum,
				Timestamp: 0,
			}

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

			res, err := mevRpcClient.SubmitBundle(context.Background(), &msgSubmitBundle)

			if err != nil {
				return err
			}

			return clientCtx.PrintProto(res)
		},
	}

	return cmd
}
