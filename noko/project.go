package noko

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

type ProjectSummary struct {
	ID               int    `json:"id"`
	Name             string `json:"name"`
	BillingIncrement int    `json:"billing_increment"`
	Enabled          bool   `json:"enabled"`
	Billable         bool   `json:"billable"`
	Color            string `json:"color"`
	URL              string `json:"url"`
}

type Project struct {
	ID                int           `json:"id"`
	Name              string        `json:"name"`
	BillingIncrement  int           `json:"billing_increment"`
	Enabled           bool          `json:"enabled"`
	Billable          bool          `json:"billable"`
	Color             string        `json:"color"`
	URL               string        `json:"url"`
	Group             Group         `json:"group"`
	Minutes           int           `json:"minutes"`
	BillableMinutes   int           `json:"billable_minutes"`
	UnbillableMinutes int           `json:"unbillable_minutes"`
	InvoicedMinutes   int           `json:"invoiced_minutes"`
	RemainingMinutes  int           `json:"remaining_minutes"`
	BudgetedMinutes   int           `json:"budgeted_minutes"`
	Import            Import        `json:"import"`
	Invoices          []Invoice     `json:"invoices"`
	Participants      []Participant `json:"participants"`
	Entries           int           `json:"entries"`
	EntriesURL        string        `json:"entries_url"`
	Expenses          int           `json:"expenses"`
	ExpensesURL       string        `json:"expenses_url"`
	CreatedAt         time.Time     `json:"created_at"`
	UpdatedAt         time.Time     `json:"updated_at"`
	MergeURL          string        `json:"merge_url"`
	ArchiveURL        string        `json:"archive_url"`
	UnarchiveURL      string        `json:"unarchive_url"`
}

func (c *Client) GetProjects(ctx context.Context) ([]Project, error) {
	req, err := http.NewRequest("GET", fmt.Sprintf("%s/projects", c.BaseURL), nil)
	if err != nil {
		return nil, err
	}

	var projects []Project
	err = c.sendRequest(ctx, req, &projects)
	if err != nil {
		return nil, err
	}

	return projects, nil
}
