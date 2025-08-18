// Package workerpool будет лежать паттерни workerpool
package workerpool

import (
	"fmt"
)

// Воркер - основное действие которое будет делатся
func (p *Pool) worker() {
	defer p.wg.Done()

	for url := range p.in {
		resp, err := p.client.Get(url)
		if err != nil {
			p.out <- fmt.Sprintf("Error: %s  url: %s", err, url)
			continue
		}
		p.out <- fmt.Sprintf("OK %s status: %s", url, resp.Status)
		resp.Body.Close()
	}
}
