package cmd

import (
	"fmt"
	"log"
	"nina/backend"
	"nina/noko"

	"github.com/spf13/cobra"
)

func NewEntryCmd(m backend.Backend) *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "entries",
		Aliases: []string{"entry"},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all entries",
		Run:   BackendRunCmd(m, listEntriesCmd),
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func listEntriesCmd(m backend.Backend) {
	entries, err := m.GetEntries()

	if err != nil {
		log.Fatal(err)
	}

	for _, entry := range entries {
		outputEntry(m, &entry)
	}
}

func outputEntry(m backend.Backend, entry *noko.Entry) {
	minutes := entry.Minutes
	hours := minutes / 60
	minutes -= hours * 60
	fmt.Fprintf(m, "%s %50s %15s %2dh%02d:  %s\n", entry.Date, entry.Project.Name, entry.User.FirstName, hours, minutes, entry.Description)
}
