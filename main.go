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

	go func() {
		for _, url := range urls {
			in <- url
		}
		close(in)
	}()

	for i := 0; i < workerCount; i++ {
		go workerpool.Worker(in, out, &wg, client)
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
