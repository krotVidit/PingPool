// Package workerpool - пакет служащий для реализации паттерна workerpool
// Создание структуры и конструктора Pool - где уже создается Пул Ворекров
package workerpool

import (
	"net/http"
	"sync"
)

type Pool struct {
	wg     sync.WaitGroup
	client *http.Client
	in     chan string
	Result chan Results
}

func NewPool(workerCount int, client *http.Client) *Pool {
	p := &Pool{
		client: client,
		in:     make(chan string),
		Result: make(chan Results),
	}

	for i := 0; i < workerCount; i++ {
		p.wg.Add(1)
		go p.worker()
	}
	return p
}
