package noko

import (
	"context"
	"fmt"
	"net/http"
)

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
