package main

import (
	"fmt"
	"net/http"
)

// Concurrency with WaitGroups Instructions
// 1. First make the program below concurrent
// 2. Then use WaitGroups to synchronize the go routines so the
// lines print out messages as expected

func fetchUrl(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Error fetching url %s: %v\n", url, err)
		return
	}

	defer resp.Body.Close()

	fmt.Printf("Fetched url %s status %s\n", url, resp.Status)
}

func main() {
	urls := []string{
		"https://example.com",
		"https://example.org",
		"https://example.net",
	}

	for _, url := range urls {
		fetchUrl(url)
	}
}
