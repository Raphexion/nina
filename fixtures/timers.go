package fixtures

import "nina/noko"

var (
	t1 = noko.Timer{
		Seconds:     4321,
		Description: "This is foo. This is bar",
		State:       "running",
		Project: noko.ProjectSummary{
			Name: "Project 1",
		},
	}

	t2 = noko.Timer{
		Seconds:     6543,
		Description: "This is fix. This is bax",
		State:       "paused",
		Project: noko.ProjectSummary{
			Name: "Project 2",
		},
	}
)

func Timers() []noko.Timer {
	return []noko.Timer{t1, t2}
}
