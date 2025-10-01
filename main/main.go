package main

import (
	"fmt"
	"link"
	"log"
	"net/http"

	"golang.org/x/net/html"
)

func main() {
	// Example web site
	url := "https://www.cartoonnetwork.com/activate/"

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Error making HTTP request %v", err)
	}
	defer resp.Body.Close()

	r := resp.Body

	doc, err := html.Parse(r)

	fmt.Println("Successfully parsed HTML. Traversing nodes:")
	link.Traversing(doc, 0)
}
