package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nina/noko"

	"github.com/schollz/closestmatch"
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
			client := noko.NewClient()

			ctx := context.Background()
			timers, err := client.GetTimers(ctx)

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
			client := noko.NewClient()

			ctx := context.Background()
			timers, err := client.GetTimers(ctx)

			if err != nil {
				log.Fatal(err)
			}

			for _, timer := range timers {
				if timer.State == "running" {
					client := noko.NewClient()

					ctx := context.Background()
					err = client.PauseTimer(ctx, &timer)

					if err != nil {
						log.Fatal(err)
					}
				}
			}
		},
	}

	startCmd := &cobra.Command{
		Use:   "start [name of timer]",
		Short: "Start a timer",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			timer, err := timerFromName(args[0])
			if err != nil {
				log.Fatal(err)
			}

			client := noko.NewClient()

			ctx := context.Background()
			err = client.StartTimer(ctx, timer)
			if err != nil {
				log.Fatal(err)
			}
		},
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pauseCmd)
	rootCmd.AddCommand(startCmd)

	return rootCmd
}

func timerFromName(name string) (*noko.Timer, error) {
	client := noko.NewClient()

	ctx := context.Background()
	timers, err := client.GetTimers(ctx)

	if err != nil {
		return nil, err
	}

	var wordsToTest []string
	for _, timer := range timers {
		wordsToTest = append(wordsToTest, timer.Project.Name)
	}
	bagSizes := []int{2}
	cm := closestmatch.New(wordsToTest, bagSizes)

	bestName := cm.Closest(name)

	for _, timer := range timers {
		if timer.Project.Name == bestName {
			return &timer, nil
		}
	}

	return nil, errors.New("unable to find a timer")
}
