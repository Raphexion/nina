package cmd

import (
	"nina/backend"
	"nina/conf"
	"nina/fixtures"
	"strings"
	"testing"
)

func TestListTimers(t *testing.T) {
	timers := fixtures.Timers()

	back := &backend.MockBackend{Timers: timers}
	back.Init()

	conf.SetBackend(back)

	args := []string{"list"}
	cmd := NewTimerCmd()
	cmd.SetArgs(args)
	cmd.Execute()

	expected := `
Project 1                       1h12,  running: This is foo. This is bar
Project 2                       1h49,   paused: This is fix. This is bax
	       `

	if strings.TrimSpace(back.Output.String()) != strings.TrimSpace(expected) {
		t.Fatalf("Incorrect output in listCmd. expected:\n%s\n. got\n%s\n", expected, back.Output.String())
	}
}
