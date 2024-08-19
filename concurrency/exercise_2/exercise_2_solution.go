package main

import (
	"fmt"
	"time"
)

// Rate Limiter Exercise Instructions:
// 1. Make the requests concurrent using only channels
// 2. Limit the requests to 1 per second

type Client struct{}

type response struct {
	url    string
	status int
}

func (c *Client) request(url string, result chan<- *response) {
	responses := map[string]*response{
		"http://some-fake.com":   &response{status: 200, url: url},
		"http://some-fake-1.com": &response{status: 200, url: url},
		"http://some-fake-2.com": &response{status: 200, url: url},
		"http://some-fake-3.com": &response{status: 200, url: url},
	}

	result <- responses[url]
}

func fetchUrls(urls []string, c chan<- *response) {
	client := &Client{}
	ticker := time.NewTicker(time.Second * 1)
	for _, url := range urls {
		<-ticker.C
		go client.request(url, c)
	}
}

func main() {
	urls := []string{
		"http://some-fake.com",
		"http://some-fake-1.com",
		"http://some-fake-2.com",
		"http://some-fake-3.com",
	}

	respChan := make(chan *response)

	fetchUrls(urls, respChan)

	for i := 0; i < len(urls); i++ {
		fmt.Println(<-respChan)
	}

	close(respChan)
}
