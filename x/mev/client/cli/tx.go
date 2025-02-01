package cli

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/sei-protocol/sei-chain/x/mev/types"
)

// GetTxCmd returns the transaction commands for the mev module
func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      fmt.Sprintf("%s transactions subcommands", types.ModuleName),
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(
		CmdSubmitBundle(),
	)

	return cmd
}

func CmdSubmitBundle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-bundle [transactions] [block-number]",
		Short: "Submit a MEV bundle",
		Args:  cobra.ExactArgs(2),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			var txs [][]byte
			if err := json.Unmarshal([]byte(args[0]), &txs); err != nil {
				return fmt.Errorf("failed to parse transactions: %w", err)
			}

			blockNum, ok := math.ParseUint64(args[1])
			if !ok {
				return fmt.Errorf("failed to parse block number: %s", args[1])
			}

			bundle := types.Bundle{
				Sender:    "",
				Txs:       txs,
				BlockNum:  blockNum,
				Timestamp: clientCtx.Height,
			}

			msg := types.NewMsgSubmitBundle(bundle)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			msgClient := types.NewMsgClient(clientCtx)

			// TODO - proper context?
			submitBundleResponse, err := msgClient.SubmitBundle(context.Background(), msg)
			if err != nil {
				return fmt.Errorf("failed to submit bundle: %w", err)
			}

			// Do we connect to client context or can configure rpc address?
			return clientCtx.PrintProto(submitBundleResponse)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
