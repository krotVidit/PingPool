package main

import (
	"fmt"
	"net/http"
	"time"

	"ping/app/workerpool"
)

// FIXME: Мне не нравится что тут нет явной записи и нет явного вида того что я делаю, мол я конечно читаю с канала, но
// FIXME: при этом я не вижу что я делаю запись в канал и как то неявно это всё идёт ... нужно подумать
func main() {
	const workerCount = 3
	client := &http.Client{Timeout: time.Second * 10}
	pool := workerpool.NewPool(workerCount, client)

	urls := []string{ // XXX: Перенести в json, чтобы от туда читало url
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
