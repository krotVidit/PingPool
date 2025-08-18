package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"

	"ping/app/workerpool"
)

func main() {
	const workerCount = 3
	wg := sync.WaitGroup{}

	client := &http.Client{Timeout: time.Second * 10}
	in := make(chan string)
	out := make(chan string)

	urls := []string{
		"https://ya.ru/",
		"https://www.google.com/",
		"https://vk.com",
		"https://netology.ru",
	}

	go workerpool.FeedUrls(urls, in)
	go workerpool.StartWorkers(workerCount, &wg, in, out, client)

	go func() {
		wg.Wait()
		close(out)
	}()

	for result := range out {
		fmt.Println(result)
	}
}
