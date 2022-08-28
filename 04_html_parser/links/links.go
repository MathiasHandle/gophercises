package links

import (
	"io"
	"strings"

	"golang.org/x/net/html"
)

// Parse parses HTML document and returns []Link parsed from it or error
func Parse(r io.Reader) ([]Link, error) {
	doc, err := html.Parse(r)
	if err != nil {
		return nil, err
	}

	nodes := getLinkNodes(doc)

	var links []Link
	for _, node := range nodes {
		links = append(links, buildLink(node))
	}

	return links, nil
}

// Finds <a> in root node
func getLinkNodes(n *html.Node) []*html.Node {
	if n.Type == html.ElementNode && n.Data == "a" {
		return []*html.Node{n}
	}

	var ret []*html.Node
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret = append(ret, getLinkNodes(c)...)
	}

	return ret
}

// Link represents <a href="..."> in a HTML document
type Link struct {
	Href string
	Text string
}

// Builds Link from passed in *html.Node
func buildLink(n *html.Node) Link {
	var ret Link

	for _, attr := range n.Attr {
		if attr.Key == "href" {
			ret.Href = attr.Val
			break
		}
	}

	ret.Text = getText(n)
	//ret.Text = "sssssss"

	return ret
}

// returns text from inside a node, ommiting comments
func getText(n *html.Node) string {
	if n.Type == html.TextNode {
		return n.Data
	}

	if n.Type != html.ElementNode {
		return ""
	}

	var ret string
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret += getText(c) + " "
	}

	return strings.Join(strings.Fields(ret), " ")
}
