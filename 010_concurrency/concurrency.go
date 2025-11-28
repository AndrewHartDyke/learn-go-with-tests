package main

type WebsiteChecker func(string) bool
type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultsChannel := make(chan result)
	for _, url := range urls {
		go func() {
			resultsChannel <- result{url, wc(url)}
		}()
	}

	for range urls {
		r := <-resultsChannel
		results[r.string] = r.bool
	}

	return results
}
