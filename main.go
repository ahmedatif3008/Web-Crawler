package main

import (
    "fmt"
    "net/http"

    "golang.org/x/net/html"
)

func seeding(url string){
    resp, err := http.Get(url)
    if err != nil {
        fmt.Println("Error:", err)
        return
    }
    defer resp.Body.Close()

    doc, err := html.Parse(resp.Body)
    if err != nil {
        fmt.Println("Parse error:", err)
        return
    }

    var walk func(*html.Node)
    walk = func(n *html.Node) {
        if n.Type == html.ElementNode && n.Data == "a" {
            for _, attr := range n.Attr {
                if attr.Key == "href" {
                    fmt.Println(attr.Val)
                }
            }
        }

        for child := n.FirstChild; child != nil; child = child.NextSibling {
            walk(child)
        }
    }

    walk(doc)
}

func main() {
    seeding("https://pkg.go.dev/golang.org/x/net/html")
}