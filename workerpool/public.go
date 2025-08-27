// Package workerpool - пакет служащий для реализации паттерна workerpool
// Внешние функции для вызова в main
// Запись в задач в канал, выдача канала с результатом, ожидание выполнение горутин
package workerpool

import (
	"encoding/json"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (p *Pool) WriteChanIn(urls []string, interval time.Duration) {
	for {
		for _, url := range urls {
			p.in <- url
		}
		time.Sleep(interval)
		if p.stopped {
			return
		}
	}
}

func (p *Pool) ResOutChan() <-chan Result {
	return p.out
}

func (p *Pool) Wait() {
	p.wg.Wait()
	close(p.out)
}

func LoadUrls(filename string) ([]string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return nil, fmt.Errorf("не найден файл %s: %w", filename, err)
	}

	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		return nil, fmt.Errorf("неверный формат файла %s: %w", filename, err)
	}
	return urls, nil
}

func GracefulShutdown(pool *Pool) {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit
	fmt.Println("[INFO] Signal stopped...stop pool")
	pool.Stop()
	fmt.Println("[EXIT] Pool stopped corect")
}

func (p *Pool) Stop() {
	p.mu.Lock()
	defer p.mu.Unlock()

	if p.stopped {
		return
	}

	p.stopped = true
	close(p.in)
	p.wg.Wait()
}
