package cmd

import (
	"fmt"
	"log"
	"nina/backend"
	"nina/noko"
	"time"

	"github.com/spf13/cobra"
)

func NewWorkdayCmd() *cobra.Command {
	rootCmd := &cobra.Command{
		Use:   "workday",
		Short: "Summarize your workday",
		Run:   BackendRunCmd(workdayCmd),
	}

	return rootCmd
}

func workdayCmd(back backend.Backend) {
	entries, err := back.GetMyEntries()
	if err != nil {
		log.Fatal(err)
	}

	timers, err := back.GetTimers()
	if err != nil {
		log.Fatal(err)
	}

	currentTime := time.Now()
	today := currentTime.Format("2006-01-02")

	total_minutes := 0

	for _, entry := range entries {
		if entry.Date == today {

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

var workdayFormatString = "%-50s %2dh%02d  %s\n"

func outputWorkdayEntry(m backend.Backend, entry *noko.Entry) {
	minutes := entry.Minutes
	hours := minutes / 60
	minutes -= hours * 60
	fmt.Fprintf(m, workdayFormatString, entry.Project.Name, hours, minutes, entry.Description)
}

func outputWorkdayTimer(m backend.Backend, timer *noko.Timer) {
	minutes := timer.Seconds / 60
	hours := minutes / 60
	minutes -= hours * 60

	fmt.Fprintf(m, workdayFormatString, timer.Project.Name, hours, minutes, timer.Description)
}

func outputWorkdayTotal(m backend.Backend, minutes int) {
	hours := minutes / 60
	minutes -= hours * 60

	fmt.Fprintf(m, "\n")
	fmt.Fprintf(m, workdayFormatString, "", hours, minutes, "")
}
