package main

import (
	"fmt"
    "sync"
	"net/http"
	"net/url"
	"hash/fnv"

	"golang.org/x/net/html"
)

type Queue struct{
    totalElements int;
    elements []string;
    mu sync.Mutex;
    number int;
}

func (q *Queue) enqueue(url string) {
    q.mu.Lock()
    defer q.mu.Unlock()

    q.totalElements++
    q.number++
    q.elements = append(q.elements, url)
}


func (q * Queue) dequeue() string{
    q.mu.Lock()
    defer q.mu.Lock()

    q.totalElements--
    q.number--
    url := q.elements[0]
    q.elements = q.elements[1:]

    return url
}

type VisitedSet struct {
    // functions are going to be add, contains, hash
    data map[uint64]struct{};
    number int;
    mu sync.Mutex;
}

func (set * VisitedSet) add(url string){
	set.mu.Lock()
	defer set.mu.Unlock()

	set.data[hashing(url)] = struct{}{}
	set.number++
    
}

func (set * VisitedSet) contains(url string) bool{
	set.mu.Lock()
	defer set.mu.Unlock()
	hashed := hashing(url)

	_, exists :=  set.data[hashed]
	
	if !exists {
		set.data[hashed] = struct{}{}
	}

    return exists
}

func hashing(url string) uint64{
    h := fnv.New64a()
	h.Write([]byte(url))
	return h.Sum64()
}


func seeding(baseURL string) {
	resp, err := http.Get(baseURL)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		fmt.Println("HTTP error:", resp.Status)
		return
	}

	extractedURL, err := url.Parse(resp.Request.URL.String())
	if err != nil {
		fmt.Println("URL parse error:", err)
		return
	}

	doc, err := html.Parse(resp.Body)
	if err != nil {
		fmt.Println("Parse error:", err)
		return
	}

	var walk func(*html.Node)
	walk = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, attr := range n.Attr {
				if attr.Key != "href" {
					continue
				}

				ref, err := url.Parse(attr.Val)
				if err != nil {
					continue
				}

				resolved := extractedURL.ResolveReference(ref)

				// Only crawl HTTP and HTTPS URLs.
				if resolved.Scheme != "http" &&
					resolved.Scheme != "https" {
					continue
				}

				// Remove fragment identifiers.
				resolved.Fragment = ""
				resolved.RawFragment = ""

				fmt.Println(resolved.String())
			}
		}

		for child := n.FirstChild; child != nil; child = child.NextSibling {
			walk(child)
		}
	}

	walk(doc)
}


func main() {
    //seeding("https://pkg.go.dev/golang.org/x/net/html")
    fmt.Println(hashing("test poop"))
}