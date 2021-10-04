package cmd

import (
	"fmt"
	"log"
	"nina/backend"
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
		Run:   BackendRunCmd(listProjectsCmd),
	}

	rootCmd.AddCommand(listCmd)

	return rootCmd
}

func listProjectsCmd(m backend.Backend) {
	projects, err := m.GetProjects()

	if err != nil {
		log.Fatal(err)
	}

	for _, project := range projects {
		outputProject(m, &project)
	}
}

func outputProject(m backend.Backend, project *noko.Project) {
	fmt.Fprintf(m, "%-50s\n", project.Name)
}
