package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/useloopso/transaction-relayer/config"
)

var rootCmd = &cobra.Command{
	Use:   "tx-relay",
	Short: "Entrypoint for the transaction relayer",
	Args:  cobra.NoArgs,
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		if err := config.Load(); err != nil {
			fmt.Println("Failed to load config: ", err)
			os.Exit(1)
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println("Failed to start transaction relayer: ", err)
		os.Exit(1)
	}
}
