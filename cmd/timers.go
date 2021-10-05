package cmd

import (
	"errors"
	"fmt"
	"log"
	"nina/backend"
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
		Run:   BackendRunCmd(listCmdFunc),
	}

	pauseCmd := &cobra.Command{
		Use:   "pause",
		Short: "Pause active project timer",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(pauseCmdFunc),
	}

	unpauseCmd := &cobra.Command{
		Use:   "unpause",
		Short: "Unpause a paused project timer",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(unpauseCmdFunc),
	}

	noteCmd := &cobra.Command{
		Use:   "note text",
		Short: "Append a note to a timer",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(noteCmdFunc),
	}

	createCmd := &cobra.Command{
		Use:   "create",
		Short: "Create a timer for a project",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(createCmdFunc),
	}

	logCmd := &cobra.Command{
		Use:   "log",
		Short: "Log and finish timer for a given project",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(logCmdFunc),
	}

	deleteCmd := &cobra.Command{
		Use:   "delete",
		Short: "Delete a timer for a project",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(deleteCmdFunc),
	}

	adjustCmd := &cobra.Command{
		Use:   "adjust",
		Short: "Adjust the time for a project",
		Args:  cobra.NoArgs,
		Run:   BackendRunCmd(adjustCmdFunc),
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

func listCmdFunc(m backend.Backend) {
	timers, err := m.GetTimers()

	if err != nil {
		log.Fatal(err)
	}

	for _, timer := range timers {
		outputTimer(m, &timer)
	}
}

func pauseCmdFunc(m backend.Backend) {
	timer, err := m.GetRunningTimer()

	if err != nil {
		log.Fatal(err)
	}

	if err = m.PauseTimer(timer); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(m, "Paused %s\n", timer.Project.Name)
}

func unpauseCmdFunc(m backend.Backend) {
	timers, err := m.GetTimersWithState("paused")
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to unpause", timers)
	if err != nil {
		log.Fatal(err)
	}

	err = m.PauseRunningTimer()
	if err != nil {
		log.Fatal(err)
	}

	if err = m.StartTimer(timer); err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(m, "Unpaused %s\n", timer.Project.Name)
}

func noteCmdFunc(m backend.Backend) {
	timers, err := m.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to note", timers)
	if err != nil {
		log.Fatal(err)
	}

	description, err := tui.RunInput("Note", timer.Description, "I worked on ...")
	if err != nil {
		log.Fatal(err)
	}

	err = m.SetDescription(timer, description)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(m, "New note: %s\n", description)
}

func createCmdFunc(m backend.Backend) {
	projects, err := m.GetSomeProjects(false)
	if err != nil {
		log.Fatal(err)
	}

	project, err := selectAProject("Pick a project", projects)
	if err != nil {
		log.Fatal(err)
	}

	err = m.PauseRunningTimer()
	if err != nil {
		log.Fatal(err)
	}

	_, err = m.CreateTimer(project)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Fprintf(m, "Created and started timer for project: %s\n", project.Name)
}

func logCmdFunc(m backend.Backend) {
	timers, err := m.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to log", timers)
	if err != nil {
		log.Fatal(err)
	}

	if err = m.LogTimer(timer); err != nil {
		log.Fatal(err)
	}

	timer.State = "finished"
	outputTimer(m, timer)
}

func deleteCmdFunc(m backend.Backend) {
	timers, err := m.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	timer, err := selectATimer("Pick a timer to delete", timers)
	if err != nil {
		log.Fatal(err)
	}

	doDelete := promptForConfirmation(fmt.Sprintf("Are you sure you want to delete %s", timer.Project.Name))

	if !doDelete {
		return
	}

	if err = m.DeleteTimer(timer); err != nil {
		log.Fatal(err)
	}

	timer.State = "deleted"
	outputTimer(m, timer)
}

func adjustCmdFunc(m backend.Backend) {
	timers, err := m.GetTimers()
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

	err = addOrSubMinutesOnTimer(m, timer, minutes)
	if err != nil {
		log.Fatal(err)
	}

}

func addOrSubMinutesOnTimer(m backend.Backend, timer *noko.Timer, minutes int) error {
	err := m.AddOrSubTimer(timer, minutes)
	if err != nil {
		return err
	}

	return nil
}

func outputTimer(m backend.Backend, timer *noko.Timer) {
	minutes := timer.Seconds / 60
	hours := minutes / 60
	minutes -= hours * 60

	fmt.Fprintf(m, "%-30s %2dh%02d, %8s: %s\n", timer.Project.Name, hours, minutes, timer.State, timer.Description)
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
