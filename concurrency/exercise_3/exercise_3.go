package main

// Concurrency with WaitGroups Instructions
// 1. First make the program below concurrent
// 2. Then use WaitGroups to synchronize the go routines so the
// lines print out messages as expected

func fetchUrls() {
}

func main() {
	urls := []string{
		"example.com",
		"example.org",
		"example.net",
	}

	fetchUrls()
}
