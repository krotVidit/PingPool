// Package workerpool будет лежать паттерни workerpool
package workerpool

import (
	"fmt"
	"net/http"
	"sync"
)

type Pool struct {
	wg     sync.WaitGroup
	client *http.Client
	in     chan string
	out    chan string
}

func NewPool(workerCount int, client *http.Client) *Pool {
	p := &Pool{
		// Почему тут нет wg
		client: client,
		in:     make(chan string),
		out:    make(chan string),
	}

	for i := 0; i < workerCount; i++ {
		p.wg.Add(1)
		go p.worker()
	}
	return p
}

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

func (p *Pool) WriteEnqueue(urls []string) {
	go func() {
		for _, url := range urls {
			p.in <- url
		}
		close(p.in)
	}()
}

func (p *Pool) Results() <-chan string {
	return p.out
}

func (p *Pool) Wait() {
	go func() {
		p.wg.Wait()
		close(p.out)
	}()
}
