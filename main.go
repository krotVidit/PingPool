package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func main() {
	const workerCount = 3
	client := &http.Client{Timeout: time.Second * 5}
	wg := sync.WaitGroup{}
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
		wg.Add(1)
		go worker(in, out, &wg, client)
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	for resutl := range out {
		fmt.Println(resutl)
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

