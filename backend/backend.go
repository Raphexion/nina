package backend

import (
	"nina/noko"
)

// Backend
type Backend interface {
	Init() error

	// Timers
	GetTimers() ([]noko.Timer, error)
	GetTimersWithState(state string) ([]noko.Timer, error)
	GetRunningTimer() (*noko.Timer, error)
	PauseTimer(timer *noko.Timer) error
	StartTimer(timer *noko.Timer) error
	LogTimer(timer *noko.Timer) error
	CreateTimer(project *noko.Project) (*noko.Timer, error)
	DeleteTimer(timer *noko.Timer) error
	SetDescription(timer *noko.Timer, description string) error
	SetDescriptionOnRunningTimer(description string) error
	AddOrSubTimer(timer *noko.Timer, minutes int) error
	PauseRunningTimer() error
	GetProjects() ([]noko.Project, error)
	GetSomeProjects(withTimer bool) ([]noko.Project, error)

	// Entries
	GetEntries() ([]noko.Entry, error)
	GetMyEntries() ([]noko.Entry, error)

	// Output
	Write(p []byte) (n int, err error)
}
