package cmd

import (
	"fmt"
	"log"
	"nina/mid"
	"nina/noko"
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
		Run:   listCmdFunc,
	}

	pauseCmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause active timer",
		Run:   pauseCmdFunc,
	}

	unpauseCmd := &cobra.Command{
		Use:   "unpause [name of timer]",
		Short: "Unpause a paused timer",
		Args:  cobra.MinimumNArgs(1),
		Run:   unpauseCmdFunc,
	}

	noteCmd := &cobra.Command{
		Use:   "note text",
		Short: "Append a note for the running timer",
		Run:   noteCmdFunc,
	}

	createCmd := &cobra.Command{
		Use:   "create [name of project]",
		Short: "Create a timer for a project",
		Args:  cobra.MinimumNArgs(1),
		Run:   createCmdFunc,
	}

	logCmd := &cobra.Command{
		Use:   "log [name of the project]",
		Short: "Log and finish timer for a given project",
		Args:  cobra.MinimumNArgs(1),
		Run:   logCmdFunc,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete [name of the project]",
		Short: "Delete a timer",
		Args:  cobra.MinimumNArgs(1),
		Run:   deleteCmdFunc,
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pauseCmd)
	rootCmd.AddCommand(unpauseCmd)
	rootCmd.AddCommand(noteCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(deleteCmd)

	return rootCmd
}

func listCmdFunc(cmd *cobra.Command, args []string) {
	timers, err := mid.GetTimers()

	if err != nil {
		log.Fatal(err)
	}

	for _, timer := range timers {
		outputTimer(&timer)
	}
}

func pauseCmdFunc(cmd *cobra.Command, args []string) {
	timer, err := mid.GetRunningTimer()

	if err != nil {
		log.Fatal(err)
	}

	if err = mid.PauseTimer(timer); err != nil {
		log.Fatal(err)
	}

	outputTimerWithName(timer.Project.Name)
}

func unpauseCmdFunc(cmd *cobra.Command, args []string) {
	name := strings.Join(args, "")
	timer, err := mid.TimerWithName(name)
	if err != nil {
		log.Fatal(err)
	}

	if err = mid.StartTimer(timer); err != nil {
		log.Fatal(err)
	}

	outputTimerWithName(timer.Project.Name)
}

func noteCmdFunc(cmd *cobra.Command, args []string) {
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

	outputTimerWithName(timer.Project.Name)
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	name := strings.Join(args, " ")
	timer, err := mid.CreateTimer(name)
	if err != nil {
		log.Fatal(err)
	}

	outputTimerWithName(timer.Project.Name)
}

func logCmdFunc(cmd *cobra.Command, args []string) {
	name := strings.Join(args, "")
	timer, err := mid.TimerWithName(name)
	if err != nil {
		log.Fatal(err)
	}

	if err = mid.LogTimer(timer); err != nil {
		log.Fatal(err)
	}

	timer.State = "finished"
	outputTimer(timer)
}

func deleteCmdFunc(cmd *cobra.Command, args []string) {
	name := strings.Join(args, " ")
	timer, err := mid.TimerWithName(name)
	if err != nil {
		log.Fatal(err)
	}

	doDelete := promptForConfirmation(fmt.Sprintf("Are you sure you want to delete %s", timer.Project.Name))

	if !doDelete {
		return
	}

	if err = mid.DeleteTimer(timer); err != nil {
		log.Fatal(err)
	}

	timer.State = "deleted"
	outputTimer(timer)
}

// outputTimerWithName retreives the latest state of the timer
// before outputing it.
func outputTimerWithName(name string) {
	timer, err := mid.TimerWithName(name)
	if err != nil {
		log.Fatal(err)
	}
	outputTimer(timer)
}

func outputTimer(timer *noko.Timer) {
	minutes := timer.Seconds / 60
	hours := minutes / 60
	minutes -= hours * 60
	fmt.Printf("%-30s %2dh%2d, %8s: %s\n", timer.Project.Name, hours, minutes, timer.State, timer.Description)
}
