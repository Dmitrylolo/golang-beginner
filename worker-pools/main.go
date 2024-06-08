package main

import (
	"fmt"
	"log"
	"net/http"
)

type Site struct {
	URL string
}

type Result struct {
	URL    string
	Status int
}

func crawl(id int, jobs <-chan Site, result chan<- Result) {
	for site := range jobs {
		log.Printf("Worker ID: %d\n", id)
		resp, err := http.Get(site.URL)
		if err != nil {
			log.Printf(err.Error())
		}
		result <- Result{
			Status: resp.StatusCode,
			URL:    site.URL,
		}
	}
}

func main() {
	fmt.Println("Worker pools")

	jobs := make(chan Site, 4)
	results := make(chan Result, 4)

	for i := 1; i <= 4; i++ {
		go crawl(i, jobs, results)
	}

	urls := []string{"http://www.google.com", "http://x.com", "http://github.com", "http://www.facebook.com"}
	for _, url := range urls {
		jobs <- Site{URL: url}
	}
	close(jobs)

	for r := 1; r <= len(urls); r++ {
		result := <-results
		log.Println(result)
	}
}
