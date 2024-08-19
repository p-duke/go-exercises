package main

// Exercise Instructions:
// 1. Change the CheckWebsites to make concurrent requests
// 2. Run the tests to make sure they're still passing

type WebsiteChecker func(string) bool

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}


