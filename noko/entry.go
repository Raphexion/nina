package noko

import (
	"context"
	"fmt"
	"net/http"
	"time"
)

// Entry represents a Noko time entry
type Entry struct {
	ID                       int       `json:"id"`
	Date                     string    `json:"date"`
	User                     User      `json:"user"`
	Billable                 bool      `json:"billable"`
	Minutes                  int       `json:"minutes"`
	Description              string    `json:"description"`
	Project                  Project   `json:"project"`
	Tags                     []Tag     `json:"tags"`
	SourceURL                string    `json:"source_url"`
	InvoicedAt               time.Time `json:"invoiced_at"`
	Invoice                  Invoice   `json:"invoice"`
	Import                   Import    `json:"import"`
	ApprovedAt               time.Time `json:"approved_at"`
	ApprovedBy               User      `json:"approved_by"`
	URL                      string    `json:"url"`
	InvoicedOutsideOfNokoURL string    `json:"invoiced_outside_of_noko_url"`
	ApprovedURL              string    `json:"approved_url"`
	UnapprovedURL            string    `json:"unapproved_url"`
	CreatedAt                time.Time `json:"created_at"`
	UpdatedAt                time.Time `json:"updated_at"`
}

// GetEntries will get all entries
func (c *Client) GetEntries(ctx context.Context, currentUser bool) ([]Entry, error) {
	url := fmt.Sprintf("%s/entries", c.BaseURL)
	if currentUser {
		url = fmt.Sprintf("%s/current_user/entries", c.BaseURL)
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	var entries []Entry
	err = c.sendRequest(ctx, req, &entries)
	if err != nil {
		return nil, err
	}

	return entries, nil
}
