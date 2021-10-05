package backend

import (
	"context"
	"errors"
	"nina/noko"
	"os"
	"time"
)

type RealBackend struct {
	client *noko.Client
	output *os.File
}

func (m *RealBackend) Init() error {
	m.client = noko.NewClient()
	m.output = os.Stdout

	return nil
}

func (m *RealBackend) Write(p []byte) (n int, err error) {
	return m.output.Write(p)
}

func (m *RealBackend) GetTimers() ([]noko.Timer, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.GetTimers(ctx)
}

func (m *RealBackend) GetTimersWithState(state string) ([]noko.Timer, error) {
	allTimers, err := m.GetTimers()
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

func (m *RealBackend) GetRunningTimer() (*noko.Timer, error) {
	timers, err := m.GetTimers()
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

func (m *RealBackend) PauseTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.PauseTimer(ctx, timer)
}

func (m *RealBackend) StartTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.StartTimer(ctx, timer)
}

func (m *RealBackend) LogTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.LogTimer(ctx, timer)
}

func (m *RealBackend) CreateTimer(project *noko.Project) (*noko.Timer, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()

	return client.CreateTimerForProject(ctx, project)
}

func (m *RealBackend) DeleteTimer(timer *noko.Timer) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.DeleteTimer(ctx, timer)
}

func (m *RealBackend) SetDescription(timer *noko.Timer, description string) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.EditTimer(ctx, timer, description)
}

func (m *RealBackend) SetDescriptionOnRunningTimer(description string) error {
	timer, err := m.GetRunningTimer()
	if err != nil {
		return err
	}

	return m.SetDescription(timer, description)
}

func (m *RealBackend) AddOrSubTimer(timer *noko.Timer, minutes int) error {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.AddOrSubTimer(ctx, timer, minutes)
}

func (m *RealBackend) PauseRunningTimer() error {
	timer, _ := m.GetRunningTimer()
	if timer != nil {
		err := m.PauseTimer(timer)
		if err != nil {
			return err
		}
	}

	return nil
}

func (m *RealBackend) GetProjects() ([]noko.Project, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.GetProjects(ctx)
}

func (m *RealBackend) GetSomeProjects(withTimer bool) ([]noko.Project, error) {
	allProjects, err := m.GetProjects()
	if err != nil {
		return nil, err
	}

	timers, err := m.GetTimers()
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

func (m *RealBackend) GetEntries() ([]noko.Entry, error) {
	client := noko.NewClient()
	ctx, cancel := standardContext()
	defer cancel()
	return client.GetEntries(ctx, false)
}

func standardContext() (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), time.Duration(5*time.Second))
}
