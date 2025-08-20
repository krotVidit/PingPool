// Package workerpool - пакет служащий для реализации паттерна workerpool
// Results - хранит в себе результаты ответа запроса
// Report формирует удобный вывод результатов с помощью методов DurationString и StatusString
package workerpool

import (
	"fmt"
	"time"
)

type Results struct {
	URL      string
	Status   string
	Duration time.Duration
	Err      error
}

func newResults(URL, Status string, Duration time.Duration, Err error) *Results {
	return &Results{
		URL:      URL,
		Status:   Status,
		Duration: Duration,
		Err:      Err,
	}
}

func (r Results) Report() string {
	if r.Err != nil {
		return fmt.Sprintf("%-8s %-35s %v", r.StatusString(), r.URL, r.Err)
	}

	return fmt.Sprintf("%-8s %-35s %-5s %s", r.StatusString(), r.URL, r.Status, r.DurationString())
}

func (r Results) DurationString() string {
	if r.Duration < time.Second {
		return fmt.Sprintf("%dms", r.Duration.Milliseconds())
	}
	return fmt.Sprintf("%.2fs", r.Duration.Seconds())
}

func (r Results) StatusString() string {
	if r.Err != nil {
		return "[ERROR]"
	}
	return "[OK]"
}
