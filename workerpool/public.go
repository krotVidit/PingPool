// Package workerpool - пакет служащий для реализации паттерна workerpool
// Внешние функции для вызова в main
// Запись в задач в канал, выдача канала с результатом, ожидание выполнение горутин
package workerpool

import (
	"encoding/json"
	"fmt"
	"os"
)

func (p *Pool) WriteChanIn(urls []string) {
	go func() {
		for _, url := range urls {
			p.in <- url
		}
		close(p.in)
	}()
}

func (p *Pool) ResOutChan() <-chan Result {
	return p.out
}

func (p *Pool) Wait() {
	go func() {
		p.wg.Wait()
		close(p.out)
	}()
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
