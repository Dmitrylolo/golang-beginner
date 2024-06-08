package main

import (
	"fmt"
	"net/http"
	"sync"
)

var urls = []string{"https://google.com", "https://x.com", "https://github.com"}

func fetch(url string, wg *sync.WaitGroup) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(resp.Status)
	wg.Done()
}

func crawl() {
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go fetch(url, &wg)
	}
	wg.Wait()
	fmt.Println("Done")
}

func main() {
	crawl()
}
