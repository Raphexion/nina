package cmd

import (
	"log"
	"nina/backend"
	"time"

	"github.com/spf13/cobra"
)

func NewWorkweekCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "workweek",
		Short: "Summarize your workweek",
		Run:   BackendRunCmd(workweekCmd),
	}

	return rootCmd
}

func workweekCmd(back backend.Backend) {
	entries, err := back.GetMyEntries()
	if err != nil {
		log.Fatal(err)
	}

	timers, err := back.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()
	mondayTime := Monday(currentTime)
	sundayTime := Sunday(currentTime)

	total_minutes := 0

	for _, entry := range entries {
		entryDate, err := time.Parse("2006-01-02", entry.Date)
		if err != nil {
			log.Fatal(err)
		}

		if entryDate.After(mondayTime) && entryDate.Before(sundayTime) {
			outputWorkdayEntry(back, &entry)
			total_minutes += entry.Minutes
		}
	}

	for _, timer := range timers {
		outputWorkdayTimer(back, &timer)
		total_minutes += timer.Seconds / 60
	}

	outputWorkdayTotal(back, total_minutes)
}

func Monday(currentTime time.Time) time.Time {
	weekday := time.Duration(currentTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := currentTime.Date()
	currentZeroDay := time.Date(year, month, day, 0, 0, 0, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 1) * 24 * time.Hour)
}

func Sunday(currentTime time.Time) time.Time {
	weekday := time.Duration(currentTime.Weekday())
	if weekday == 0 {
		weekday = 7
	}
	year, month, day := currentTime.Date()
	currentZeroDay := time.Date(year, month, day, 23, 59, 59, 0, time.Local)
	return currentZeroDay.Add(-1 * (weekday - 7) * 24 * time.Hour)
}
