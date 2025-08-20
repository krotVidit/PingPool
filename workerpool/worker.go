// Package workerpool - пакет служащий для реализации паттерна workerpool
package workerpool

import "time"

func (p *Pool) worker() {
	defer p.wg.Done()

	for url := range p.in {
		p.Result <- *p.fetchURL(url)
	}
}

func (p *Pool) fetchURL(url string) *Results {
	start := time.Now()
	resp, err := p.client.Get(url)
	duration := time.Since(start)

	statusCode := ""
	if err == nil && resp != nil {
		statusCode = resp.Status
		resp.Body.Close()
	}
	return newResults(url, statusCode, duration, err)
}
