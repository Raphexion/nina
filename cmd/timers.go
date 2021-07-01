package cmd

import (
	"fmt"
	"log"
	"nina/mid"
	"strings"

	"github.com/spf13/cobra"
)

func NewTimerCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "timers",
		Aliases: []string{"timer"},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all timers",
		Run: func(cmd *cobra.Command, args []string) {
			timers, err := mid.GetTimers()

			if err != nil {
				log.Fatal(err)
			}

			for _, timer := range timers {
				minutes := timer.Seconds / 60
				fmt.Printf("%-30s %2d minutes, %s\n", timer.Project.Name, minutes, timer.State)
			}
		},
	}

	pauseCmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause active timer",
		Run: func(cmd *cobra.Command, args []string) {
			timer, err := mid.GetRunningTimer()

			if err != nil {
				log.Fatal(err)
			}

			if err = mid.PauseTimer(timer); err != nil {
				log.Fatal(err)
			}
		},
	}

	unpauseCmd := &cobra.Command{
		Use:   "unpause [name of timer]",
		Short: "Unpause a paused timer",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			timer, err := mid.TimerWithName(args[0])
			if err != nil {
				log.Fatal(err)
			}

			if err = mid.StartTimer(timer); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Started timer for project %s\n", timer.Project.Name)
		},
	}

	noteCmd := &cobra.Command{
		Use:   "note text",
		Short: "Append a note for the running timer",
		Run: func(cmd *cobra.Command, args []string) {
			timer, err := mid.GetRunningTimer()
			if err != nil {
				log.Fatal(err)
			}

			var text string
			if timer.Description == "" {
				text = strings.Join(args, " ")
			} else {
				text = timer.Description + ". " + strings.Join(args, " ")
			}

			if err = mid.SetDescription(text); err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Updated note for project %s\n", timer.Project.Name)
		},
	}

	createCmd := &cobra.Command{
		Use:   "create [name of project]",
		Short: "Create a timer for a project",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			timer, err := mid.CreateTimer(args[0])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("Timer created for project %s\n", timer.Project.Name)
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pauseCmd)
	rootCmd.AddCommand(unpauseCmd)
	rootCmd.AddCommand(noteCmd)
	rootCmd.AddCommand(createCmd)

	return rootCmd
}
