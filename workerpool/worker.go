// Package workerpool будет лежать паттерни workerpool
package workerpool

import (
	"fmt"
	"net/http"
	"sync"
)

func Worker(in <-chan string, out chan<- string, wg *sync.WaitGroup, client *http.Client) {
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
