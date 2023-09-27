package keywords

import (
	"github.com/Godyu97/vege9/vegeTools"
	"golang.org/x/net/html"
	"io"
)

func ParseHtml(r io.Reader) (string, error) {
	n, err := html.Parse(r)
	if err != nil {
		return "", err
	}
	return formatHTML(n), nil
}

var ArrBlock = []string{
	"h1", "h2", "h3", "h4", "h5",
	"p",
	"br", "b",
	"div",
	"td",
}

func formatHTML(n *html.Node) string {
	var result string
	if n.Type == html.ElementNode && vegeTools.ItemIsInSlice(n.Data, ArrBlock) {
		result += "\n"
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.ElementNode:
			result += formatHTML(c)
		case html.TextNode:
			result += c.Data
		}
	}
	return result
}
