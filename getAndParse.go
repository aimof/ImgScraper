package ImgScraper

import (
	"io"
	"log"
	"net/http"
	"net/url"

	"golang.org/x/net/html"
)

func getURLsPic(rowURL string) []string {
	urlsPic := getAndParse(rowURL)
	checkedURLsPic := checkHost(rowURL, urlsPic)
	return checkedURLsPic
}

func getAndParse(rowURL string) (urlsPic []string) {
	results := getPage(rowURL)
	for _, result := range results {
		urlsPic = append(urlsPic, result)
	}
	return urlsPic
}

func checkHost(rowURL string, urlsPic []string) []string {
	parsedURL, err := url.Parse(rowURL)
	if err != nil {
		log.Fatal(err)
	}
	host := parsedURL.Host
	var newUrlsPic []string
	slash := "/"
	for _, urlPic := range urlsPic {
		if urlPic[0] == slash[0] {
			urlPic = "http://" + host + urlPic
		}
		newUrlsPic = append(newUrlsPic, urlPic)

	}
	return newUrlsPic
}

func parseItem(r io.Reader) []string {
	var results []string
	doc, err := html.Parse(r)
	if err != nil {
		log.Fatal(err)
	}

	var result string
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "img" {
			for _, img := range n.Attr {
				if img.Key == "src" {
					result = img.Val
					results = append(results, result)
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	return results
}

func getPage(rowURL string) []string {
	res, err := http.Get(rowURL)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	results := parseItem(res.Body)
	return results
}
