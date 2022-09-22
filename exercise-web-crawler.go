package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, checker *Checker) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	
	done := make(chan bool)
	spawned := 0
	for _, u := range urls {
		u := u
		if checker.HasVisited(u) {
			continue
		}
		// fmt.Printf("%v %p, %v\n", i, &u, u)
		spawned++
		go func() {
			Crawl(u, depth-1, fetcher, checker)
			checker.Update(u)
			done <- true
		}()
		// go func(url string) {
		// 	Crawl(url, depth-1, fetcher, checker)
		// 	checker.Update(url)
		// 	done <- true
		// }(u)
	}
	for i := 0; i < spawned; i++ {
		<-done
	}
	return
}

func main() {
	checker := &Checker{
		crawled: make(map[string]bool),
	}
	Crawl("https://golang.org/", 4, fetcher, checker)
}

type Checker struct {
	mu      sync.Mutex
	crawled map[string]bool
}

func (c *Checker) HasVisited(url string) bool {
	c.mu.Lock()
	defer c.mu.Unlock()
	if _, ok := c.crawled[url]; ok {
		return true
	}
	return false
}

func (c *Checker) Update(url string) {
	c.mu.Lock()
	c.crawled[url] = true
	c.mu.Unlock()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher struct {
	mu    sync.Mutex
	cache map[string]*fakeResult
}

type fakeResult struct {
	body string
	urls []string
}

func (f *fakeFetcher) Fetch(url string) (string, []string, error) {
	f.mu.Lock()
	defer f.mu.Unlock()
	if res, ok := f.cache[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = &fakeFetcher{
	cache: map[string]*fakeResult{
		"https://golang.org/": &fakeResult{
			"The Go Programming Language",
			[]string{
				"https://golang.org/pkg/",
				"https://golang.org/cmd/",
			},
		},
		"https://golang.org/pkg/": &fakeResult{
			"Packages",
			[]string{
				"https://golang.org/",
				"https://golang.org/cmd/",
				"https://golang.org/pkg/fmt/",
				"https://golang.org/pkg/os/",
			},
		},
		"https://golang.org/pkg/fmt/": &fakeResult{
			"Package fmt",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
		"https://golang.org/pkg/os/": &fakeResult{
			"Package os",
			[]string{
				"https://golang.org/",
				"https://golang.org/pkg/",
			},
		},
	},
}
