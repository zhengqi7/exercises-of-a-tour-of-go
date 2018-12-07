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
 
type SafeMap struct {
    mp  map[string]int
    mux sync.Mutex
}
 
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl ( sm SafeMap, url string, depth int, fetcher Fetcher) {
    defer wg.Done()
	sm.mux.Lock()
    sm.mp[url]++
    if 1 < sm.mp[url] {
        sm.mux.Unlock()
        return
    }
    sm.mux.Unlock()
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
    for _, u := range urls {
		wg.Add(1)
        go Crawl(sm, u, depth-1, fetcher)
    }
    return
}

var wg sync.WaitGroup
func main() {
	wg.Add(1)
    go Crawl( SafeMap{mp: make(map[string]int)}, "http://golang.org/", 4, fetcher)
	wg.Wait()
}
 
// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult
 
type fakeResult struct {
    body string
    urls []string
}
 
func (f fakeFetcher) Fetch(url string) (string, []string, error) {
    if res, ok := f[url]; ok {
        return res.body, res.urls, nil
    }
    return "", nil, fmt.Errorf("not found: %s", url)
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