package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uuid-cli",
	Short: "A tool for UUID generation and parsing",
	Run:   help,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func help(cmd *cobra.Command, args []string) {
	cmd.Help()
}
