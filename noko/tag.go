package noko

type Tag struct {
	ID            int    `json:"id"`
	Name          string `json:"name"`
	Billable      bool   `json:"billable"`
	FormattedName string `json:"formatted_name"`
	URL           string `json:"url"`
}
