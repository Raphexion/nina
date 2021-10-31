package cmd

import (
	"testing"
	"time"
)

func TestWorkweekMonday(t *testing.T) {
	tests := []struct {
		dateString       string
		dateComment      string
		mondayDateString string
		sundayDateString string
	}{
		{"2011-01-19 01:02:03", "Wednesday", "2011-01-17 00:00:00", "2011-01-23 23:59:59"},
		{"2011-01-17 01:02:03", "Monday", "2011-01-17 00:00:00", "2011-01-23 23:59:59"},
		{"2011-01-23 01:02:03", "Sunday", "2011-01-17 00:00:00", "2011-01-23 23:59:59"},
	}

	templateDate := "2006-01-02 15:04:05"

	for _, tt := range tests {
		d, err := time.Parse(templateDate, tt.dateString)
		if err != nil {
			t.Fatalf("Unable to parse date string %s", tt.dateString)
		}

		m, err := time.Parse(templateDate, tt.mondayDateString)
		if err != nil {
			t.Fatalf("Unable to parse date string %s", tt.mondayDateString)
		}

		s, err := time.Parse(templateDate, tt.sundayDateString)
		if err != nil {
			t.Fatalf("Unable to parse date string %s", tt.sundayDateString)
		}

		if !equalDate(Monday(d), m) {
			t.Fatalf("Incorrect Monday for %s. got: %s. expeced: %s", tt.dateString, Monday(d), tt.mondayDateString)
		}

		if !equalDate(Sunday(d), s) {
			t.Fatalf("Incorrect Sunday for %s. got: %s. expeced: %s", tt.dateString, Sunday(d), tt.sundayDateString)
		}
	}

}

func equalDate(t1, t2 time.Time) bool {
	return t1.Year() == t2.Year() && t1.YearDay() == t2.YearDay()
}
