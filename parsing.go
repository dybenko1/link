package link

import (
	"io"
	"log"
	"net/http"
	"strings"

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

// Traverse through all nodes of the website and return the parsed links
func TraverseAndParseLinks(n *html.Node, parsedLinks []Link) []Link {
	// Base Case: If we found a <a> link then we traverse to get all the text
	isLinkElement := n.Type == html.ElementNode && n.Data == "a"
	if isLinkElement {
		var url string
		var text string
		// Getting URL
		for _, attribute := range n.Attr {
			if attribute.Key == "href" {
				url = attribute.Val
				break
			}
		}
		// Getting text  and creating link element for this url-text
		text = textFinder(n, "")
		parsedLinks = append(parsedLinks, Link{
			url,
			text,
		})
	}

	// Recursive, first deep (because we get to the son and then the
	// son of the son, after the last son we get to the sibling of
	// the upper layer and so on).
	// If we have identified a LinkHTMLElement we do not need to go deeper, but to its sibling
	if isLinkElement {
		for c := n.NextSibling; c != nil; c = c.NextSibling {
			parsedLinks = TraverseAndParseLinks(c, parsedLinks)
		}
	} else {
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			parsedLinks = TraverseAndParseLinks(c, parsedLinks)
		}
	}
	return parsedLinks
}

// Once we find a link element we traverse throughout all its text elements and return a concatenation
func textFinder(n *html.Node, concatText string) string {
	// If we find a text element we concatenate it to the current string variable
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data) // Avoids blank space, also "\n"
		if text != "" {
			concatText = concatText + " " + text
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		concatText = textFinder(c, concatText)
	}
	return concatText
}

type Link struct {
	Href string
	Text string
}
