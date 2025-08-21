package main

import (
	"fmt"
	"net/http"
	"time"

	"ping/app/workerpool"
)

func main() {
	const workerCount = 3
	client := &http.Client{Timeout: time.Second * 10}
	pool := workerpool.NewPool(workerCount, client)

	urls := []string{
		"https://ya.ru/",
		"https://www.google.com/",
		"https://vk.com",
		"https://netology.ru",
	}

	pool.WriteChanIn(urls)
	pool.Wait()

	for result := range pool.ResultsOutChan() {
		fmt.Println(result.Report())
	}
}
