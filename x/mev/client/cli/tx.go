package cli

import (
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
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

			var txs []string
			if err := json.Unmarshal([]byte(args[0]), &txs); err != nil {
				return fmt.Errorf("failed to parse transactions: %w", err)
			}

			blockNum, err := ParseUint64(args[1])
			if err != nil {
				return fmt.Errorf("failed to parse block number: %w", err)
			}

			bundle := types.Bundle{
				Sender:    clientCtx.GetFromAddress().String(),
				Txs:       txs,
				BlockNum:  blockNum,
				Timestamp: clientCtx.Height,
			}

			msg := types.NewMsgSubmitBundle(bundle)
			if err := msg.ValidateBasic(); err != nil {
				return err
			}

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
