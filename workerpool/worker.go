// Package workerpool будет лежать паттерни workerpool
package workerpool

import "time"

// Воркер - основное действие которое будет делатся
func (p *Pool) worker() {
	defer p.wg.Done()

	for url := range p.in {
		start := time.Now()
		resp, err := p.client.Get(url)
		duration := time.Since(start)

		result := Result{
			URL:      url,
			Duration: duration,
			Error:    err,
		}

		if err == nil && resp != nil {
			result.Status = resp.Status
			resp.Body.Close()
		}
		p.Result <- result
	}
}
