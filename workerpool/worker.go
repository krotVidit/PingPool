// Package workerpool - пакет служащий для реализации паттерна workerpool
// timeTrack - может принять в себя любуфую функцию и посчитать время её выполнения
// worker - получает результаты запроса из handleURL и отправляет их структуру Results
package workerpool

import "time"

func (p *Pool) worker() {
	defer p.wg.Done()

	for url := range p.in {
		p.out <- p.handleURL(url)
	}
}

func timeTrack[T any](f func() (T, error)) (res T, duration time.Duration, err error) {
	start := time.Now()
	res, err = f()
	return res, time.Since(start), err
}

func (p *Pool) handleURL(url string) Result {
	status, duration, err := timeTrack(func() (string, error) {
		resp, err := p.client.Get(url)
		if err != nil {
			return "", err
		}
		defer resp.Body.Close()
		return resp.Status, nil
	})
	return newResult(url, status, duration, err)
}
