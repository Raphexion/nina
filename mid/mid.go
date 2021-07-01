package mid

import (
	"context"
	"errors"
	"nina/noko"

	"github.com/schollz/closestmatch"
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

func TimerWithName(name string) (*noko.Timer, error) {
	timers, err := GetTimers()

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

func SetDescription(description string) error {
	timer, err := GetRunningTimer()
	if err != nil {
		return err
	}

	client := noko.NewClient()
	ctx := context.Background()
	return client.EditTimer(ctx, timer, description)
}
