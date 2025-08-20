package workerpool

func (p *Pool) WriteEnqueue(urls []string) {
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
