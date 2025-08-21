// Package workerpool - пакет служащий для реализации паттерна workerpool
// Внешние функции для вызова в main
// Запись в задач в канал, выдача канала с результатом, ожидание выполнение горутин
package workerpool

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
