package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	const workerCount = 3
	wg := sync.WaitGroup{}

	client := &http.Client{Timeout: time.Second * 5}
	in := make(chan string)
	out := make(chan string)

	urls := []string{
		"https://ya.ru/",
		"https://www.google.com/",
		"https://vk.com",
		"https://netology.ru",
	}

	go func() {
		for _, url := range urls {
			in <- url
		}
		close(in)
	}()

	for i := 0; i < workerCount; i++ {
		go worker(in, out, &wg, client)
		wg.Add(1)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for result := range out {
		fmt.Println(result)
	}
}

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

