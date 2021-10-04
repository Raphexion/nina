package fixtures

import "nina/noko"

var (
	p1 = noko.Project{Name: "Project 1"}

	p2 = noko.Project{Name: "Project 2"}
)

func Projects() []noko.Project {
	return []noko.Project{p1, p2}
}
