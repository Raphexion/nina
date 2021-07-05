package cmd

import (
	"context"
	"fmt"
	"log"
	"nina/noko"

	"github.com/spf13/cobra"
)

func NewEntryCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "entries",
		Aliases: []string{"entry"},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all entries",
		Run: func(cmd *cobra.Command, args []string) {
			client := noko.NewClient()

			ctx := context.Background()
			entries, err := client.GetEntries(ctx, false)

			if err != nil {
				log.Fatal(err)
			}

			for _, entry := range entries {
				fmt.Printf("%-10s %-80s %-10d\n", entry.User.FirstName, entry.Description, entry.Minutes)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}
