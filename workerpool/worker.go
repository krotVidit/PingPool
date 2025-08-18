// Package workerpool будет лежать паттерни workerpool
package workerpool

import (
	"fmt"
	"net/http"
	"sync"
)

func worker(in <-chan string, out chan<- string, wg *sync.WaitGroup, client *http.Client) {
	defer wg.Done()

	for url := range in {
		resp, err := client.Get(url)
		if err != nil {
			out <- fmt.Sprintf("Error: %s  url: %s", err, url)
			continue
		}
		out <- fmt.Sprintf("OK %s status: %s", url, resp.Status)
		resp.Body.Close()
	}
}

func FeedUrls(urls []string, in chan<- string) {
	go func() {
		for _, url := range urls {
			in <- url
		}
		close(in)
	}()
}

func StartWorkers(workerCount int, wg *sync.WaitGroup, in <-chan string, out chan<- string, client *http.Client) {
	for i := 0; i < workerCount; i++ {
		go worker(in, out, wg, client)
		wg.Add(1)
	}
}
