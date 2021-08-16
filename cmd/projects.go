package cmd

import (
	"context"
	"fmt"
	"log"
	"nina/noko"

	"github.com/spf13/cobra"
)

func NewProjectCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "projects",
		Aliases: []string{"project"},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all projects",
		Run: func(cmd *cobra.Command, args []string) {
			client := noko.NewClient()

			ctx := context.Background()
			projects, err := client.GetProjects(ctx)

			if err != nil {
				log.Fatal(err)
			}

			for _, project := range projects {
				fmt.Printf("%-50s\n", project.Name)
			}
		},
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}
