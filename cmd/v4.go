package cmd

import (
	"fmt"
	"os"

	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

var v4Options = struct{
	ShowsDetails bool
}{}

var v4Cmd = &cobra.Command{
	Use:   "v4 [count]",
	Short: "Generates a new UUID in version 4 format",
	Long: fmt.Sprintf("Generates the specified count of UUIDs in version 4 format. The count must be a positive integer and must not exceed %d.", MAX_COUNT),
	Args: cobra.MaximumNArgs(1),
	Run: generateV4,
}

func init() {
	rootCmd.AddCommand(v4Cmd)

	v4Cmd.Flags().BoolVarP(&v4Options.ShowsDetails, "details", "d", false, "Display detailed output")
}

func generateV4(cmd *cobra.Command, args []string) {
	count, err := getCountFromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for i := 0; i < count; i++ {
		id, err := uuid.NewV4()
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error Generating UUID: %v\n", err)
			os.Exit(1)
		}

		if v4Options.ShowsDetails {
			if err := printDetails(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			fmt.Println(id.String())
		}
	}
}
