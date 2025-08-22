package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"ping/app/workerpool"
)

func main() {
	const workerCount = 3
	client := &http.Client{Timeout: time.Second * 10}
	pool := workerpool.NewPool(workerCount, client)

	data, err := os.ReadFile("urls.json")
	if err != nil {
		fmt.Println("Ошибка - не найден файл urls.json")
		return
	}

	var urls []string
	if err := json.Unmarshal(data, &urls); err != nil {
		fmt.Println("Ошибка - неверный формат urls.json")
		return
	}

	pool.WriteChanIn(urls)
	pool.Wait()

	for result := range pool.ResOutChan() {
		fmt.Println(result.Report())
	}
}
