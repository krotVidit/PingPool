// Package workerpool будет лежать паттерни workerpool
package workerpool

import "time"

// Воркер - основное действие которое будет делатся
func (p *Pool) worker() {
	defer p.wg.Done()

	for url := range p.in {
		result := p.fetchURL(url)
		p.Result <- result
	}
}

func (p *Pool) fetchURL(url string) Results {
	start := time.Now()
	resp, err := p.client.Get(url)
	duration := time.Since(start)

	result := Results{
		URL:      url,
		Duration: duration,
		Err:      err,
	}

	if err == nil && resp != nil {
		result.Status = resp.Status
		resp.Body.Close()
	}
	return result
}
