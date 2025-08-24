package main

import (
	"fmt"
	"net/http"
	"time"

	"ping/app/workerpool"
)

// TODO: В данной вертке будет реаилзована возможность вызывать  в цикле пинг
func main() {
	const workerCount = 3
	const interval = time.Second * 10
	client := &http.Client{Timeout: time.Second * 10}
	pool := workerpool.NewPool(workerCount, client)

	urls, err := workerpool.LoadUrls("urls.json")
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	pool.WriteChanIn(urls, interval)
	pool.Wait()

	for result := range pool.ResOutChan() {
		fmt.Println(result.Report())
	}

	go workerpool.GracefulShutdown(pool)
}
