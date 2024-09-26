package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/araddon/dateparse"
	"github.com/gofrs/uuid/v5"
	"github.com/spf13/cobra"
)

var v7Options = struct{
	ShowsDetails bool
	Timestamp string
}{}

var v7Cmd = &cobra.Command{
	Use:   "v7 [count]",
	Short: "Generates a new UUID in version 7 format",
	Long:  fmt.Sprintf("Generates the specified count of UUIDs in version 7 format. The count must be a positive integer and must not exceed %d.", MAX_COUNT),
	Args:  cobra.MaximumNArgs(1),
	Run:   generateV7,
}

func init() {
	rootCmd.AddCommand(v7Cmd)

	v7Cmd.Flags().BoolVarP(&v7Options.ShowsDetails, "details", "d", false, "Display detailed output")
	v7Cmd.Flags().StringVarP(&v7Options.Timestamp, "timestamp", "t", "", "Timestamp to parse (optional)")
}

func generateV7(cmd *cobra.Command, args []string) {
	count, err := getCountFromArgs(args)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	var t time.Time

	for i := 0; i < count; i++ {
		var id uuid.UUID

		if v7Options.Timestamp != "" {
			t, err = dateparse.ParseAny(v7Options.Timestamp)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error: Failed to parse timestamp: %v\n", err)
				os.Exit(1)
			}

			id, err = uuid.NewV7AtTime(t)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error Generating UUID: %v\n", err)
				os.Exit(1)
			}
		} else {
			id, err = uuid.NewV7()
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error Generating UUID: %v\n", err)
				os.Exit(1)
			}
		}

		if v7Options.ShowsDetails {
			if err := printV7Details(id); err != nil {
				fmt.Fprintf(os.Stderr, "Error: %v\n", err)
			}
		} else {
			fmt.Println(id.String())
		}
	}
}
