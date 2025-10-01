package link

import (
	"fmt"
	"strings"

	"golang.org/x/net/html"
)

func Traversing(n *html.Node, depth int) {
	// Indent to display each level depending on the depth of the Node in the Tree
	indent := strings.Repeat("	", depth)

	// We identify the type of Node.
	// They could be an Element (E.g. <p>, <div>),(it includes its Attributes (a property of an element) in "".Attr
	// When getting attributes we get a Slice, because there could be multiple
	// Text (What is inside an Element), etc.

	// Part 1: HTML Element
	if n.Type == html.ElementNode {
		// The data of an Element Node is the tag, E.g. "p", out of "<p>"
		fmt.Printf("%s<%s>\n", indent, n.Data)
	}

	// Part 2: Text of HTML Element
	if n.Type == html.TextNode {
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%sText: %s\n", indent, text)
		}
	}

	// Recursive Loop
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Traversing(c, depth+1)
	}

}

// Recursive function to traverse the node tree
func Traverse_2(n *html.Node, depth int) {
	// Print indentation for three visualization
	indent := strings.Repeat("	", depth)

	// Check if the node is an Element (like <html>, <h1>, <p>)
	if n.Type == html.ElementNode {
		fmt.Printf("%s<%s>\n", indent, n.Data)
	}

	// Check if the node is Text and print its content
	if n.Type == html.TextNode {
		// Only print not empty text nodes after trimming whitespace
		text := strings.TrimSpace(n.Data)
		if text != "" {
			fmt.Printf("%s(Text: %s)\n", indent, text)
		}
	}

	// Recurse on children
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		Traverse_2(c, depth+1)
	}

}
