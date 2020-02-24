package crawler

import (
	"golang.org/x/net/html"
	"net/http"
)

// Parse an html response and return all valid urls from a tags
func parseHtml(response *http.Response) []string {

	defer response.Body.Close()

	var links []string

	z := html.NewTokenizer(response.Body)

	for {
		tt := z.Next()

		switch tt {
			case html.ErrorToken:
				return links
			case html.StartTagToken, html.EndTagToken:
				token := z.Token()
				if token.Data == "a" {
					for _, attr := range token.Attr {
						if attr.Key == "href" {
							links = append(links, attr.Val)
						}
					}
				}

		}
	}
}
