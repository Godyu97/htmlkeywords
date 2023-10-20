package keywords

import (
	"github.com/Godyu97/vege9/vege"
	"golang.org/x/net/html"
	"io"
	"strings"
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
	"ul", "li", "ol", "dl", "dt", "dd",
}

func formatHTML(n *html.Node) string {
	var result string
	if n.Type == html.ElementNode && vege.ItemIsInSlice(n.Data, ArrBlock) {
		result += "\n"
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		switch c.Type {
		case html.ElementNode:
			if n.Data == "table" {
				//处理表格
				result += formatTable(c)
			} else {
				result += formatHTML(c)
			}
		case html.TextNode:
			result += c.Data
		}
	}
	return result
}

/*
tr td1 td2 td3
tr td11 td22 td33
tr td111 td222 td333

-->

	td1
	td11
	td111

	td2
	td22
	td222

	td3
	td33
	td333
*/
func formatTable(n *html.Node) string {
	t := make([][]string, 0)
	tContent := "\n"
	//解析table
	for tr := n.FirstChild; tr != nil; tr = tr.NextSibling {
		switch tr.Type {
		case html.ElementNode:
			if tr.Data == "tr" {
				trLine := make([]string, 0)
				for td := tr.FirstChild; td != nil; td = td.NextSibling {
					tdText := formatHTML(td)
					tdText = strings.TrimSpace(tdText)
					if tdText != "" {
						trLine = append(trLine, tdText)
					}
				}
				t = append(t, trLine)
			} else {
				tContent += formatTable(tr)
			}
		case html.TextNode:
			tContent += tr.Data
		}
	}
	if len(t) > 0 {
		ltr := len(t[0])
		if ltr == 2 {
			// 横向表格
			for _, columns := range t {
				for i, _ := range columns {
					tContent += columns[i]
					tContent += "\n"
				}
			}
		} else {
			//竖向表格
			for i := 0; i < ltr; i++ {
				for _, columns := range t {
					if i < len(columns) {
						tContent += columns[i]
						tContent += "\n"
					}
				}
			}
		}
	}
	return tContent
}
