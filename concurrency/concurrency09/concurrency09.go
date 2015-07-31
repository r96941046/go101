package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of url and a slice of urls found on that page
	Fetch(url string) (body string, urls []string, err error)
}

func _Crawl(url string, fetcher Fetcher, urlsChannel chan []string) {

	body, urls, err := fetcher.Fetch(url)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Found: %s %q\n", url, body)
	}

	urlsChannel <- urls
}

// Crawl uses fetcher to recursively crawl pages starting with url, to a maximum depth
func Crawl(url string, depth int, fetcher Fetcher, history map[string]bool) {

	if depth <= 0 {
		return
	}

	urlsChannel := make(chan []string)

	history[url] = true
	urlToFetch := 1

	go _Crawl(url, fetcher, urlsChannel)

	for urlToFetch > 0 {
		urlToFetch--
		urls := <-urlsChannel
		for _, url := range urls {
			if _, done := history[url]; !done {
				history[url] = true
				urlToFetch++
				go _Crawl(url, fetcher, urlsChannel)
			}
		}
	}
}

func main() {
	history := make(map[string]bool)
	Crawl("http://golang.org/", 4, fetcher, history)
}

// fakeFetcher is Fetcher that returns canned results
type fakeFetcher map[string]*fakeResult

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

type fakeResult struct {
	body string
	urls []string
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"http://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"http://golang.org/pkg/",
			"http://golang.org/cmd/",
		},
	},
	"http://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"http://golang.org/",
			"http://golang.org/cmd/",
			"http://golang.org/pkg/fmt/",
			"http://golang.org/pkg/os/",
		},
	},
	"http://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
	"http://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"http://golang.org/",
			"http://golang.org/pkg/",
		},
	},
}
