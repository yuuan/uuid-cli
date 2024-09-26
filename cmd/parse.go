package cmd

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

var parseCmd = &cobra.Command{
	Use:   "parse [UUID]",
	Short: "Parse the given UUID and displays its details",
	Long: "Parse the given UUID and displays its details",
	Args: cobra.MinimumNArgs(1),
	Run: parse,
}

func init() {
	rootCmd.AddCommand(parseCmd)
}

func parse(cmd *cobra.Command, args []string) {
	for _, arg := range args {
		id, err := uuid.FromString(arg)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Parse Error: %v", err)
		}

		if id.Version() == 7 {
			if err := printV7Details(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			if err := printDetails(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		}
	}
}
