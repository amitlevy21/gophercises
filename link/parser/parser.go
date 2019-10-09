package parser

import (
	"io"

	"golang.org/x/net/html"
)

// Link holds data about links in a html file
type Link struct {
	Href string
	Text string
}

// Parse parses links from html input
func Parse(in io.Reader) []Link {
	z := html.NewTokenizer(in)
	return loopTokens(z)
}

func loopTokens(z *html.Tokenizer) (links []Link) {
	for {
		tt := z.Next()
		if tt == html.ErrorToken {
			break
		}
		for {
			k, v, b := z.TagAttr()
			if string(k) == "href" {
				href := string(v)
				z.Next()
				text := string(z.Text())
				links = append(links, Link{Href: href, Text: text})
			}
			if !b {
				break
			}
		}
	}
	return links
}
