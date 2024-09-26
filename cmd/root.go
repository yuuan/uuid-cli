package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "uuid-cli",
	Short: "Generates a new UUID in version 7 format",
	Long:  fmt.Sprintf("Generates the specified count of UUIDs in version 7 format. The count must be a positive integer and must not exceed %d.", MAX_COUNT),
	Args:  cobra.MaximumNArgs(1),
	Run:   generateV7,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {}
