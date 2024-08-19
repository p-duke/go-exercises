package main

import "log"

// Rate Limiter Exercise Instructions:
// 1. Make the requests concurrent using only channels
// 2. Limit the requests to 1 per second

type Client struct{}

type response struct {
	status int
}

func (c *Client) request(url string) *response {
	responses := map[string]*response{
		"http://some-fake.com":   &response{status: 200},
		"http://some-fake-1.com": &response{status: 200},
		"http://some-fake-2.com": &response{status: 200},
		"http://some-fake-3.com": &response{status: 200},
	}

	return responses[url]
}

func fetchUrls(urls []string) {
	client := &Client{}
	for _, url := range urls {
		response := client.request(url)

		if response.status != 200 {
			log.Println("Request failed")
		} else {
			log.Printf("Request to %s successful", url)
		}
	}
}

func main() {
	urls := []string{
		"http://some-fake.com",
		"http://some-fake-1.com",
		"http://some-fake-2.com",
		"http://some-fake-3.com",
	}

	fetchUrls(urls)
}
