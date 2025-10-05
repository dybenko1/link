package main

import (
	"fmt"
	"io"
	"link"
	"log"
	"os"

	"golang.org/x/net/html"
)

func main() {
	// Example web site
	//url := "https://www.cartoonnetwork.com/activate/"
	//request := link.Get_website(url)

	file, err := os.Open("examples/ex1.html")
	if err != nil {
		log.Fatal("Error opening file: %v", err)
	}

	defer file.Close()

	var ex_site io.Reader = file

	doc, err := html.Parse(ex_site)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully parsed HTML. Traversing nodes:")

	link.LinkParser(doc)
}
