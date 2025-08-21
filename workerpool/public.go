// Package workerpool - пакет служащий для реализации паттерна workerpool
// Внешние функции для вызова в main
// Запись в задач в канал, выдача канала с результатом, ожидание выполнение горутин
package workerpool

func (p *Pool) WriteChanIn(urls []string) { // XXX: Можно наверное в Pool перенести ?
	go func() {
		for _, url := range urls {
			p.in <- url
		}
		close(p.in)
	}()
}

func (p *Pool) ResultsOutChan() <-chan Results {
	return p.Result
}

func (p *Pool) Wait() {
	go func() {
		p.wg.Wait()
		close(p.Result)
	}()
}
