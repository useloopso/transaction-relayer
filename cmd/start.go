package cmd

import (
	"github.com/spf13/cobra"
	"github.com/useloopso/transaction-relayer/core"
)

var startCmd = &cobra.Command{
	Use:   "start",
	Short: "Starts the transaction relayer",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		core.RunRouter()
	},
}

func init() {
	rootCmd.AddCommand(startCmd)
}
