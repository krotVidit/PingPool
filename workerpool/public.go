package workerpool

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
