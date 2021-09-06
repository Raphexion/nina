package noko

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Timer represents a Noko time struct and is one of the most important
// structures in Noko and Nina.
type Timer struct {
	ID                   int            `json:"id"`
	State                string         `json:"state"`
	Seconds              int            `json:"seconds"`
	FormattedTime        string         `json:"formatted_time"`
	Date                 string         `json:"date"`
	Description          string         `json:"description"`
	User                 User           `json:"user"`
	Project              ProjectSummary `json:"project"`
	URL                  string         `json:"url"`
	StartURL             string         `json:"start_url"`
	PauseURL             string         `json:"pause_url"`
	AddOrSubtractTimeURL string         `json:"add_or_subtract_time_url"`
	LogURL               string         `json:"log_url"`
}

// GetTimers will get all timers from Noko
func (c *Client) GetTimers(ctx context.Context) ([]Timer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/timers", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var timers []Timer
	err = c.sendRequest(ctx, req, &timers)
	if err != nil {
		return nil, err
	}

	return timers, nil
}

// GetTimer will get a specific timer from Noko
func (c *Client) GetTimer(ctx context.Context, id int) (*Timer, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/timers/%d", c.BaseURL, id), nil)
	if err != nil {
		return nil, err
	}

	timer := &Timer{}
	err = c.sendRequest(ctx, req, timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// StartTimer will start a Noko timer
func (c *Client) StartTimer(ctx context.Context, timer *Timer) error {
	req, err := http.NewRequest("PUT", timer.StartURL, nil)
	if err != nil {
		return err
	}

	err = c.send(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// PauseTimer will start a Noko timer
func (c *Client) PauseTimer(ctx context.Context, timer *Timer) error {
	req, err := http.NewRequest("PUT", timer.PauseURL, nil)
	if err != nil {
		return err
	}

	err = c.send(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// LogTimer will log/finish a Noko timer and register the time
func (c *Client) LogTimer(ctx context.Context, timer *Timer) error {
	req, err := http.NewRequest("PUT", timer.LogURL, nil)
	if err != nil {
		return err
	}

	err = c.send(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// EditTimer will change the description of a timer.
func (c *Client) EditTimer(ctx context.Context, timer *Timer, description string) error {
	values := map[string]string{"description": description}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("PUT", timer.URL, bytes.NewBuffer(jsonValue))
	if err != nil {
		return err
	}

	err = c.send(ctx, req)
	if err != nil {
		return err
	}

	return nil
}

// CreateTimerForProject will create and start a new timer for a project
func (c *Client) CreateTimerForProject(ctx context.Context, project *Project) (*Timer, error) {
	req, err := http.NewRequest("PUT", fmt.Sprintf("%s/timer/start", project.URL), nil)
	if err != nil {
		return nil, err
	}

	timer := &Timer{}
	err = c.sendRequest(ctx, req, timer)
	if err != nil {
		return nil, err
	}

	return timer, nil
}

// DeleteTimer will delete a time without loggin the time.
func (c *Client) DeleteTimer(ctx context.Context, timer *Timer) error {
	project := timer.Project
	req, err := http.NewRequest("DELETE", fmt.Sprintf("%s/timer", project.URL), nil)

	if err != nil {
		return err
	}

	return c.send(ctx, req)
}

// AddOrSubTimer will add or subtract minutes from project timer
func (c *Client) AddOrSubTimer(ctx context.Context, timer *Timer, minutes int) error {
	values := map[string]int{"minutes": minutes}
	jsonValue, _ := json.Marshal(values)

	req, err := http.NewRequest("PUT", timer.AddOrSubtractTimeURL, bytes.NewBuffer(jsonValue))

	if err != nil {
		return err
	}

	err = c.send(ctx, req)
	if err != nil {
		return err
	}

	return nil
}
