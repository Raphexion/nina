package backend

import (
	"bytes"
	"errors"
	"nina/noko"
)

var ErrNotImplemented = errors.New("not implemented")

type MockBackend struct {
	Projects  []noko.Project
	Timers    []noko.Timer
	Entries   []noko.Entry
	MyEntries []noko.Entry
	Output    *bytes.Buffer
}

func (m *MockBackend) Init() error {
	b := make([]byte, 0, 4096)
	m.Output = bytes.NewBuffer(b)
	return nil
}

func (m *MockBackend) Write(p []byte) (n int, err error) {
	return m.Output.Write(p)
}

func (m *MockBackend) GetTimers() ([]noko.Timer, error) {
	return m.Timers, nil
}

func (m *MockBackend) GetTimersWithState(state string) ([]noko.Timer, error) {
	return nil, ErrNotImplemented
}

func (m *MockBackend) GetRunningTimer() (*noko.Timer, error) {
	return nil, ErrNotImplemented
}

func (m *MockBackend) PauseTimer(timer *noko.Timer) error {
	return ErrNotImplemented
}
func (m *MockBackend) StartTimer(timer *noko.Timer) error {
	return ErrNotImplemented
}

func (m *MockBackend) LogTimer(timer *noko.Timer) error {
	return ErrNotImplemented
}

func (m *MockBackend) CreateTimer(project *noko.Project) (*noko.Timer, error) {
	return nil, ErrNotImplemented
}

func (m *MockBackend) DeleteTimer(timer *noko.Timer) error {
	return ErrNotImplemented
}

func (m *MockBackend) SetDescription(timer *noko.Timer, description string) error {
	return ErrNotImplemented
}

func (m *MockBackend) SetDescriptionOnRunningTimer(description string) error {
	return ErrNotImplemented
}

func (m *MockBackend) AddOrSubTimer(timer *noko.Timer, minutes int) error {
	return ErrNotImplemented
}

func (m *MockBackend) PauseRunningTimer() error {
	return ErrNotImplemented
}

func (m *MockBackend) GetProjects() ([]noko.Project, error) {
	return m.Projects, nil
}

func (m *MockBackend) GetSomeProjects(withTimer bool) ([]noko.Project, error) {
	return nil, ErrNotImplemented
}

func (m *MockBackend) GetEntries() ([]noko.Entry, error) {
	return m.Entries, nil
}

func (m *MockBackend) GetMyEntries() ([]noko.Entry, error) {
	return m.MyEntries, nil
}
