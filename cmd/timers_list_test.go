package cmd

import (
	"nina/backend"
	"nina/fixtures"
	"os"
	"strings"
	"testing"
)

func TestListTimers(t *testing.T) {
	os.Setenv("NOKO_ACCESS_TOKEN", "XYZ")

	timers := fixtures.Timers()

	backend := &backend.MockBackend{Timers: timers}
	backend.Init()

	args := []string{"list"}
	cmd := NewTimerCmd(backend)
	cmd.SetArgs(args)
	cmd.Execute()

	expected := `
Project 1                       1h12,  running: This is foo. This is bar
Project 2                       1h49,   paused: This is fix. This is bax
	       `

	if strings.TrimSpace(backend.Output.String()) != strings.TrimSpace(expected) {
		t.Fatalf("Incorrect output in listCmd. expected:\n%s\n. got\n%s\n", expected, backend.Output.String())
	}
}
