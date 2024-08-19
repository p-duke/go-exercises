package main

// Exercise Instructions:
// 1. Change the CheckWebsites to make concurrent requests
// 2. Run the tests to make sure they're still passing

type WebsiteChecker func(string) bool
type result struct {
	url string
	success bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.url] = r.success
	}

	return results
}



