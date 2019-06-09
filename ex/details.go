package ex

import (
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/riyadennis/realty-tool/internal"
)

func viewProperty(url, id string) *internal.PropertyRecord {
	fullURL := fmt.Sprintf("%s%s", url, id)
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("error while loading the property :: %v", err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalf("unable to view property :: %v", err)
	}
	return checkAndUpdate(id, fullURL, doc)
}

func price(doc *goquery.Document) string {
	var price string
	doc.Find("#price_overlay").Each(func(arg1 int, elm *goquery.Selection) {
		price = elm.Text()
	})
	return strings.TrimSpace(price)
}

func status(doc *goquery.Document) string {
	var status string
	doc.Find(".overlay-text").Each(func(arg1 int, elm *goquery.Selection) {
		status = strings.Replace(elm.Text(), "\n", "", -1)
	})
	return strings.TrimSpace(status)
}

func name(doc *goquery.Document) string {
	var name string
	doc.Find("#property_heading").Each(func(arg1 int, elm *goquery.Selection) {
		h1 := elm.Find("h1")
		name = strings.Replace(h1.Text(), "\n", "", -1)
	})
	return strings.Trim(name, "")
}
