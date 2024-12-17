package cli

import (
	"encoding/base64"
	"strconv"
	"strings"

	"github.com/cosmos/cosmos-sdk/client"
	"github.com/cosmos/cosmos-sdk/client/flags"
	"github.com/cosmos/cosmos-sdk/client/tx"
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/spf13/cobra"

	"github.com/sei-protocol/sei-chain/x/mev/types"
)

func GetTxCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:                        types.ModuleName,
		Short:                      "MEV transaction subcommands",
		DisableFlagParsing:         true,
		SuggestionsMinimumDistance: 2,
		RunE:                       client.ValidateCmd,
	}

	cmd.AddCommand(CmdSubmitBundle())
	return cmd
}

func CmdSubmitBundle() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "submit-bundle [transactions] [block-height] [bundle-fee]",
		Short: "Submit a bundle of transactions",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			clientCtx, err := client.GetClientTxContext(cmd)
			if err != nil {
				return err
			}

			// Parse transactions (base64 encoded)
			txsStr := args[0]
			txs := make([][]byte, 0)
			for _, txStr := range strings.Split(txsStr, ",") {
				tx, err := base64.StdEncoding.DecodeString(txStr)
				if err != nil {
					return err
				}
				txs = append(txs, tx)
			}

			// Parse block height
			height, err := strconv.ParseInt(args[1], 10, 64)
			if err != nil {
				return err
			}

			// Parse bundle fee
			fee, err := sdk.ParseCoinNormalized(args[2])
			if err != nil {
				return err
			}

			msg := types.NewMsgSubmitBundle(
				clientCtx.GetFromAddress(),
				txs,
				height,
				fee,
			)

			return tx.GenerateOrBroadcastTxCLI(clientCtx, cmd.Flags(), msg)
		},
	}

	flags.AddTxFlagsToCmd(cmd)
	return cmd
}
