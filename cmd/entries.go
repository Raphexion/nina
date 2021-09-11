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
				outputEntry(&entry)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func outputEntry(entry *noko.Entry) {
	minutes := entry.Minutes
	hours := minutes / 60
	minutes -= hours * 60
	fmt.Printf("%s %50s %15s %2dh%02d:  %s\n", entry.Date, entry.Project.Name, entry.User.FirstName, hours, minutes, entry.Description)
}
