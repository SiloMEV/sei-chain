package cmd

import (
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/sei-protocol/sei-chain/mev/cli"
	tmcli "github.com/tendermint/tendermint/libs/cli"

	//nolint:gosec,G108
	"github.com/spf13/cobra"
	_ "net/http/pprof"
)

//nolint:gosec
func MEVCmd(defaultMEVRPCAddr string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mev",
		Short: "MEV commands",
		Long:  "MEV commands to operate on off-chain data",
	}

	cmd.PersistentFlags().String(mev.FlagMEVRpcAddr, defaultMEVRPCAddr, "MEV RPC address")
	cmd.PersistentFlags().StringP(tmcli.OutputFlag, "o", "text", "Output format (text|json)")

	cmd.AddCommand(cli.CmdPendingBundles())
	//cmd.AddCommand(cli.CmdSubmitBundles())

	return cmd
}
