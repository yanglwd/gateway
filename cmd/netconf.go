package cmd

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {
	netConfCmd.PersistentFlags().String("addr", "0.0.0.0:12345", "listen addr")
	netConfCmd.PersistentFlags().Uint16("id", 1000, "server id")

	viper.BindPFlag("addr", netConfCmd.PersistentFlags().Lookup("addr"))
	viper.BindPFlag("id", netConfCmd.PersistentFlags().Lookup("id"))

	rootCmd.AddCommand(netConfCmd)
}

var netConfCmd = &cobra.Command{
	Use:   "netconf",
	Short: "net config.",
	Long:  ``,
}
