package main

import (
	"fmt"
	"net/http"
	"time"

	"ping/app/workerpool"
)

func main() {
	const (
		workerCount   = 3
		pingInterval  = time.Second * 10
		clientTimeout = time.Second * 10
	)

	client := &http.Client{Timeout: clientTimeout}
	pool := workerpool.NewPool(workerCount, client)

	urls, err := workerpool.LoadUrls("urls.json")
	if err != nil {
		fmt.Println("Ошибка", err)
		return
	}

	go pool.WriteChanIn(urls, pingInterval)
	go pool.Wait()

	go workerpool.GracefulShutdown(pool)

	for result := range pool.ResOutChan() {
		fmt.Println(result.Report())
	}
}
