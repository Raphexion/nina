package mid

import (
	"context"
	"errors"
	"nina/noko"
	"nina/utils"
)

func GetTimers() ([]noko.Timer, error) {
	client := noko.NewClient()
	ctx := context.Background()
	return client.GetTimers(ctx)
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
	ctx := context.Background()
	return client.PauseTimer(ctx, timer)
}

func StartTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx := context.Background()
	return client.StartTimer(ctx, timer)
}

func LogTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx := context.Background()
	return client.LogTimer(ctx, timer)
}

func CreateTimer(projectName string) (*noko.Timer, error) {
	project, err := ProjectWithName(projectName)
	if err != nil {
		return nil, err
	}

	client := noko.NewClient()
	ctx := context.Background()
	return client.CreateTimerForProject(ctx, project)
}

func DeleteTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx := context.Background()
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

	bestName := utils.ClosestMatch(name, alternatives)

	for _, timer := range timers {
		if timer.Project.Name == bestName {
			return &timer, nil
		}
	}

	return nil, errors.New("unable to find a timer")
}

func SetDescription(description string) error {
	timer, err := GetRunningTimer()
	if err != nil {
		return err
	}

	client := noko.NewClient()
	ctx := context.Background()
	return client.EditTimer(ctx, timer, description)
}

func GetProjects() ([]noko.Project, error) {
	client := noko.NewClient()
	ctx := context.Background()
	return client.GetProjects(ctx)
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

	bestName := utils.ClosestMatch(name, alternatives)

	for _, project := range projects {
		if project.Name == bestName {
			return &project, nil
		}
	}

	return nil, errors.New("unable to find a timer")
}
