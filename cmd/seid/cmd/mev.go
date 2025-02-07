package cmd

import (
	"github.com/sei-protocol/sei-chain/mev"
	"github.com/sei-protocol/sei-chain/mev/cli"
	"github.com/spf13/cobra"

	//nolint:gosec,G108
	_ "net/http/pprof"
)

//nolint:gosec
func MEVCmd(defaultMEVRPCAddr string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "mev",
		Short: "MEV commands",
		Long:  "MEV commands to operate on off-chain data",
		//RunE: func(cmd *cobra.Command, _ []string) error {
		//
		//	serverCtx := server.GetServerContextFromCmd(cmd)
		//	if err := serverCtx.Viper.BindPFlags(cmd.Flags()); err != nil {
		//		return err
		//	}
		//
		//	//mevRpcAddr := serverCtx.Viper.GetString(flagMEVRpcAddr)
		//
		//	return nil
		//},
	}

	cmd.PersistentFlags().String(mev.FlagMEVRpcAddr, defaultMEVRPCAddr, "MEV RPC address")

	cmd.AddCommand(cli.CmdPendingBundles())
	cmd.AddCommand(cli.CmdSubmitBundles())

	return cmd
}
