package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"

	"github.com/mathiashandle/04htmlparser/links"
)

/*
	1. GET page
	2. parse all the links on page
	3. build proper urls with our links
		- filter out external urls
		- add domain for links like /about
	4. find all pages (BFS)
	5. print out XML
*/

func main() {
	urlFlag := flag.String("url", "https://gophercises.com", "url that you want to build a sitemap for")
	maxDepth := flag.Int("depth", 10, "the max number of links deep to traverse")
	flag.Parse()

	pages := bfs(*urlFlag, *maxDepth)
	for _, page := range pages {
		fmt.Println(page)
	}
}

// implementation of BFS algorithmn
func bfs(urlString string, maxDepth int) []string {
	// map of all already visited links
	seen := make(map[string]struct{})
	// queue of currently processed links
	var q map[string]struct{}
	// newQueue of links to be processed
	nq := map[string]struct{}{
		urlString: struct{}{},
	}

	for i := 0; i <= maxDepth; i++ {
		// move nq to q and create new empty nq
		q, nq = nq, make(map[string]struct{})
		for url, _ := range q {
			// if the link was already seen, skip it
			if _, ok := seen[url]; ok {
				continue
			}
			// mark link as seen
			seen[url] = struct{}{}
			// get new links from currently visited page
			for _, link := range getPages(url) {
				nq[link] = struct{}{}
			}
		}
	}

	var ret []string
	for url, _ := range seen {
		ret = append(ret, url)
	}

	return ret
}

// visits a page and returns all links from it
func getPages(urlStr string) []string {
	resp, err := http.Get(urlStr)
	if err != nil {
		fmt.Printf("error calling get request, %s", err)
	}
	defer resp.Body.Close()

	reqUrl := resp.Request.URL
	baseUrl := &url.URL{
		Scheme: reqUrl.Scheme,
		Host:   reqUrl.Host,
	}
	base := baseUrl.String()

	return filter(getHrefs(resp.Body, base), withPrefix(base))
}

// returns all links from a page
func getHrefs(r io.Reader, base string) []string {
	links, _ := links.Parse(r)

	var hrefs []string
	for _, l := range links {
		switch {
		case strings.HasPrefix(l.Href, "/"):
			hrefs = append(hrefs, base+l.Href)

		case strings.HasPrefix(l.Href, "http"):
			hrefs = append(hrefs, l.Href)
		}
	}

	return hrefs
}

// filters out links by passed in filter fn
func filter(links []string, keepFn func(string) bool) []string {
	var ret []string
	for _, link := range links {
		if keepFn(link) {
			ret = append(ret, link)
		}
	}

	return ret
}

func withPrefix(pfx string) func(string) bool {
	return func(link string) bool {
		return strings.HasPrefix(link, pfx)
	}
}
