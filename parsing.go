package link

import (
	"io"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func Get_website(url string) io.Reader {

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request %v", err)
	}

	r := resp.Body
	return r
}

func LinkParser(n *html.Node) {
	// To store link
	var ElementLink string
	var LinkText string

	// Check if it is an HTML Element of type <a>
	if n.Type == html.ElementNode && n.Data == "a" {
		// Extracting link
		for _, attr := range n.Attr {
			if attr.Key == "href" {
				ElementLink = attr.Val
			}
		}
		// Now every text or HTML element will be a child of <a>
		// So we are going to extract all text, HTML element or TextElement
		for inner_c := n.FirstChild; inner_c != nil; inner_c = inner_c.NextSibling {
			// TextElement
			if inner_c.Type == html.TextNode {
				// This text comes directly from <a>. It is a direct child of <a>
				LinkText = LinkText + inner_c.Data

			}
			// If the first child of children element es HTML element
			// That its first child is HTML element
			// This code ignore nested html that may contain text
			if inner_c.Type == html.ElementNode && inner_c.FirstChild.Type == html.TextNode {
				LinkText = LinkText + inner_c.FirstChild.Data
			}
		}

	}
	// Recursing code
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		LinkParser(c)
	}
}

type link struct {
	Href string
	Text string
}
