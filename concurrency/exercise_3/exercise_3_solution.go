package main

import (
	"fmt"
	"net/http"
	"sync"
)

// Concurrency with WaitGroups Instructions
// 1. First make the program below concurrent
// 2. Then use WaitGroups to synchronize the go routines so the
// lines print out messages as expected

func fetchUrl(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered from panic in fetchURL: %v\n", r)
		}
	}()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching url %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Fetched url %s status %s\n", url, resp.Status)
}

func main() {
	var wg sync.WaitGroup
	urls := []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
	}

	for _, url := range urls {
		wg.Add(1)
		go fetchUrl(url, &wg)
	}

	fmt.Println("Waiting for all goroutines...")
	wg.Wait()
	fmt.Println("All requests completed.")
}
