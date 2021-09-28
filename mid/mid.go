package mid

import (
	"context"
	"errors"
	"nina/noko"
	"nina/utils"
	"time"
)

func GetTimers() ([]noko.Timer, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.GetTimers(ctx)
}

func GetTimersWithState(state string) ([]noko.Timer, error) {
	allTimers, err := GetTimers()
	if err != nil {
		return nil, err
	}

	var timers []noko.Timer
	for _, timer := range allTimers {
		if timer.State == state {
			timers = append(timers, timer)
		}
	}

	return timers, nil
}

func GetRunningTimer() (*noko.Timer, error) {
	timers, err := GetTimers()
	if err != nil {
		return nil, err
	}

	for _, timer := range timers {
		if timer.State == "running" {
			return &timer, nil
		}
	}

	return nil, errors.New("no running timer found")
}

func PauseTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.PauseTimer(ctx, timer)
}

func StartTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.StartTimer(ctx, timer)
}

func LogTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.LogTimer(ctx, timer)
}

func CreateTimer(projectName string) (*noko.Timer, error) {
	project, err := ProjectWithName(projectName)
	if err != nil {
		return nil, err
	}

	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()

	return client.CreateTimerForProject(ctx, project)
}

func DeleteTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.DeleteTimer(ctx, timer)
}

func TimerWithName(name string) (*noko.Timer, error) {
	timers, err := GetTimers()

	if err != nil {
		return nil, err
	}

	var alternatives []string
	for _, timer := range timers {
		alternatives = append(alternatives, timer.Project.Name)
	}

	bestName, err := utils.ClosestMatch(name, alternatives)
	if err != nil {
		return nil, err
	}

	for _, timer := range timers {
		if timer.Project.Name == bestName {
			return &timer, nil
		}
	}

	return nil, errors.New("unable to find a timer")
}

func SetDescription(timer *noko.Timer, description string) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.EditTimer(ctx, timer, description)
}

func SetDescriptionOnRunningTimer(description string) error {
	timer, err := GetRunningTimer()
	if err != nil {
		return err
	}

	return SetDescription(timer, description)
}

func AddOrSubTimer(timer *noko.Timer, minutes int) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.AddOrSubTimer(ctx, timer, minutes)
}

func PauseRunningTimer() error {
	timer, _ := GetRunningTimer()
	if timer != nil {
		err := PauseTimer(timer)
		if err != nil {
			return err
		}
	}

	return nil
}

func GetProjects() ([]noko.Project, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.GetProjects(ctx)
}

func GetSomeProjects(withTimer bool) ([]noko.Project, error) {
	allProjects, err := GetProjects()
	if err != nil {
		return nil, err
	}

	timers, err := GetTimers()
	if err != nil {
		return nil, err
	}

	timerProjectNames := make(map[string]bool)
	for _, timer := range timers {
		timerProjectNames[timer.Project.Name] = true
	}

	var projects []noko.Project
	for _, project := range allProjects {
		if _, ok := timerProjectNames[project.Name]; ok == withTimer {
			projects = append(projects, project)
		}
	}

	return projects, nil
}

func ProjectWithName(name string) (*noko.Project, error) {
	projects, err := GetProjects()

	if err != nil {
		return nil, err
	}

	var alternatives []string
	for _, project := range projects {
		alternatives = append(alternatives, project.Name)
	}

	bestName, err := utils.ClosestMatch(name, alternatives)
	if err != nil {
		return nil, err
	}

	for _, project := range projects {
		if project.Name == bestName {
			return &project, nil
		}
	}

	return nil, errors.New("unable to find a timer")
}

func standardContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(5*time.Second))
}
