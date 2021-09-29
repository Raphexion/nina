package cmd

import (
	"errors"
	"fmt"
	"log"
	"nina/mid"
	"nina/noko"
	"nina/tui"
	"nina/utils"

	"github.com/spf13/cobra"
)

func NewTimerCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:     "timers",
		Aliases: []string{"timer"},
	}

	listCmd := &cobra.Command{
		Use:   "list",
		Short: "List all project timers",
		Args:  cobra.NoArgs,
		Run:   listCmdFunc,
	}

	pauseCmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause active project timer",
		Args:  cobra.NoArgs,
		Run:   pauseCmdFunc,
	}

	unpauseCmd := &cobra.Command{
		Use:   "unpause",
		Short: "Unpause a paused project timer",
		Args:  cobra.NoArgs,
		Run:   unpauseCmdFunc,
	}

	noteCmd := &cobra.Command{
		Use:   "note text",
		Short: "Append a note to a timer",
		Args:  cobra.NoArgs,
		Run:   noteCmdFunc,
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a timer for a project",
		Args:  cobra.NoArgs,
		Run:   createCmdFunc,
	}

	logCmd := &cobra.Command{
		Use:   "log",
		Short: "Log and finish timer for a given project",
		Args:  cobra.NoArgs,
		Run:   logCmdFunc,
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a timer for a project",
		Args:  cobra.NoArgs,
		Run:   deleteCmdFunc,
	}

	adjustCmd := &cobra.Command{
		Use:   "adjust",
		Short: "Adjust the time for a project",
		Args:  cobra.NoArgs,
		Run:   adjustCmdFunc,
	}

	rootCmd.AddCommand(listCmd)
	rootCmd.AddCommand(pauseCmd)
	rootCmd.AddCommand(unpauseCmd)
	rootCmd.AddCommand(noteCmd)
	rootCmd.AddCommand(createCmd)
	rootCmd.AddCommand(logCmd)
	rootCmd.AddCommand(deleteCmd)
	rootCmd.AddCommand(adjustCmd)

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

	fmt.Printf("Paused %s\n", timer.Project.Name)
}

func unpauseCmdFunc(cmd *cobra.Command, args []string) {
	timers, err := mid.GetTimersWithState("paused")
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to unpause", timers)
	if err != nil {
		log.Fatal(err)
	}

	err = mid.PauseRunningTimer()
	if err != nil {
		log.Fatal(err)
	}

	if err = mid.StartTimer(timer); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Unpaused %s\n", timer.Project.Name)
}

func noteCmdFunc(cmd *cobra.Command, args []string) {
	timers, err := mid.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to unpause", timers)
	if err != nil {
		log.Fatal(err)
	}

	description, err := tui.RunInput("Note", timer.Description, "I worked on ...")
	if err != nil {
		log.Fatal(err)
	}

	err = mid.SetDescription(timer, description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("New note: %s\n", description)
}

func createCmdFunc(cmd *cobra.Command, args []string) {
	projects, err := mid.GetSomeProjects(false)
	if err != nil {
		log.Fatal(err)
	}

	project, err := selectAProject("Pick a project", projects)
	if err != nil {
		log.Fatal(err)
	}

	err = mid.PauseRunningTimer()
	if err != nil {
		log.Fatal(err)
	}

	_, err = mid.CreateTimer(project.Name)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Created and started timer for project: %s\n", project.Name)
}

func logCmdFunc(cmd *cobra.Command, args []string) {
	timers, err := mid.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to log", timers)
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
	timers, err := mid.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to log", timers)
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

func adjustCmdFunc(cmd *cobra.Command, args []string) {
	timers, err := mid.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to adjust", timers)
	if err != nil {
		log.Fatal(err)
	}

	format, err := tui.RunInput("How much to adjust", "", "-1h23m")
	if err != nil {
		log.Fatal(err)
	}

	minutes, err := utils.MinutesFromHMFormat(format)
	if err != nil {
		log.Fatal(err)
	}

	err = addOrSubMinutesOnTimer(timer, minutes)
	if err != nil {
		log.Fatal(err)
	}

}

func addOrSubMinutesOnTimer(timer *noko.Timer, minutes int) error {
	err := mid.AddOrSubTimer(timer, minutes)
	if err != nil {
		return err
	}

	return nil
}

func outputTimer(timer *noko.Timer) {
	minutes := timer.Seconds / 60
	hours := minutes / 60
	minutes -= hours * 60
	fmt.Printf("%-30s %2dh%02d, %8s: %s\n", timer.Project.Name, hours, minutes, timer.State, timer.Description)
}

func selectATimer(title string, timers []noko.Timer) (*noko.Timer, error) {
	if len(timers) == 0 {
		return nil, errors.New("no timers found")
	}

	if len(timers) == 1 {
		return &timers[0], nil
	}

	choices := make([]string, len(timers))

	var timerIndex int

	for ii, timer := range timers {
		choices[ii] = timer.Project.Name
	}

	timerIndex, err := tui.RunTuiSelector(title, choices)
	if err != nil {
		return nil, err
	}

	return &timers[timerIndex], nil
}

func selectAProject(title string, projects []noko.Project) (*noko.Project, error) {
	if len(projects) == 0 {
		return nil, errors.New("no projects found")
	}

	if len(projects) == 1 {
		return &projects[0], nil
	}

	choices := make([]string, len(projects))

	var projectIndex int

	for ii, project := range projects {
		choices[ii] = project.Name
	}

	projectIndex, err := tui.RunTuiSelector(title, choices)
	if err != nil {
		return nil, err
	}

	return &projects[projectIndex], nil
}
