package cmd

import (
	"context"
	"errors"
	"fmt"
	"log"
	"nina/noko"
	"os"

	"github.com/schollz/closestmatch"
	"github.com/spf13/cobra"
)

var listTimersCmd = &cobra.Command{
	Use:   "list",
	Short: "List all timers",
	Run: func(cmd *cobra.Command, args []string) {
		client := noko.NewClient(os.Getenv("NOKO_API_KEY"))

		ctx := context.Background()
		timers, err := client.GetTimers(ctx)

		if err != nil {
			log.Fatal(err)
		}

		for _, timer := range timers {
			minutes := timer.Seconds / 60

			fmt.Printf("%-30s %2d minutes, %s\n", timer.Project.Name, minutes, timer.State)
			fmt.Printf("%s\n", timer.URL)
		}
	},
}

var startTimerCmd = &cobra.Command{
	Use:   "start [name of timer]",
	Short: "Start a timer",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timer, err := timerFromName(args[0])
		if err != nil {
			log.Fatal(err)
		}

		client := noko.NewClient(os.Getenv("NOKO_API_KEY"))

		ctx := context.Background()
		err = client.StartTimer(ctx, timer)
		if err != nil {
			log.Fatal(err)
		}
	},
}

var pauseTimerCmd = &cobra.Command{
	Use:   "pause [name of timer]",
	Short: "Pause a timer",
	Args:  cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		timer, err := timerFromName(args[0])
		if err != nil {
			log.Fatal(err)
		}

		client := noko.NewClient(os.Getenv("NOKO_API_KEY"))

		ctx := context.Background()
		err = client.PauseTimer(ctx, timer)
		if err != nil {
			log.Fatal(err)
		}
	},
}

func timerFromName(name string) (*noko.Timer, error) {
	client := noko.NewClient(os.Getenv("NOKO_API_KEY"))

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
