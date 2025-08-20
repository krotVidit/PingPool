package workerpool

import (
	"fmt"
	"time"
)

type Result struct {
	URL      string
	Status   string
	Duration time.Duration
	Error    error
}

func (r Result) Report() string {
	if r.Error != nil {
		return fmt.Sprintf("%-8s %-35s %v", r.StatusString(), r.URL, r.Error)
	}

	return fmt.Sprintf("%-8s %-35s %-5s %s", r.StatusString(), r.URL, r.Status, r.DurationString())
}

func (r Result) DurationString() string {
	if r.Duration < time.Second {
		return fmt.Sprintf("%dms", r.Duration.Milliseconds())
	}
	return fmt.Sprintf("%.2fs", r.Duration.Seconds())
}

func (r Result) StatusString() string {
	if r.Error != nil {
		return "[ERROR]"
	}
	return "[OK]"
}
