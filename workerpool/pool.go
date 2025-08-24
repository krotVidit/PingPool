// Package workerpool - пакет служащий для реализации паттерна workerpool
// Создание структуры и конструктора Pool - где уже создается Пул Ворекров
package workerpool

import (
	"net/http"
	"sync"
)

type Pool struct {
	wg          sync.WaitGroup
	client      *http.Client
	workerCount int
	stopped     bool
	mu          sync.Mutex
	in          chan string
	out         chan Result
}

func NewPool(workerCount int, client *http.Client) *Pool {
	p := &Pool{
		client:      client,      // Можно ещё создать метод по умолчанию который будет создавать клиента
		workerCount: workerCount, // Пригодится всё равно
		stopped:     false,       // по сути мы же можем реализовать один прогон или бесконечный через это
		in:          make(chan string),
		out:         make(chan Result),
	}

	for i := 0; i < workerCount; i++ {
		p.wg.Add(1)
		go p.worker()
	}
	return p
}
