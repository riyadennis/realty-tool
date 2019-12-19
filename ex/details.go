package ex

import (
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/riyadennis/realty-tool/internal"
)

func property(url, id string) *internal.PropertyRecord {
	if id == "" {
		return nil
	}
	fullURL := fmt.Sprintf("%s%s", url, id)
	resp, err := http.Get(fullURL)
	if err != nil {
		log.Fatalf("error while loading the property :: %v", err)
		return nil
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalf("unable to view property :: %v", err)
		return nil
	}
	re := regexp.MustCompile("[0-9]+")
	n := name(doc)
	p := price(doc)
	s := status(doc)
	if n == "" || p == "" || s == "" {
		return nil
	}
	return &internal.PropertyRecord{
		ID:     re.FindAllString(id, -1)[0],
		URL:    fullURL,
		Price:  price(doc),
		Status: status(doc),
		Name:   n,
	}
}

func price(doc *goquery.Document) string {
	var price string
	doc.Find("#price_overlay").Each(func(arg1 int, elm *goquery.Selection) {
		price = elm.Text()
	})
	if price == "" {
		doc.Find("#propertyHeaderPrice").Each(func(arg1 int, elm *goquery.Selection) {
			price = elm.Find("strong").Text()
		})
	}
	return strings.TrimSpace(price)
}

func status(doc *goquery.Document) string {
	var status string
	doc.Find(".overlay-text").Each(func(arg1 int, elm *goquery.Selection) {
		status = strings.Replace(elm.Text(), "\n", "", -1)
	})
	if status == "" {
		doc.Find("#propertyHeaderPrice").Each(func(arg1 int, elm *goquery.Selection) {
			status = elm.Find("small").Text()
		})
	}
	return strings.TrimSpace(status)
}

func name(doc *goquery.Document) string {
	var name string
	doc.Find("#property_heading").Each(func(arg1 int, elm *goquery.Selection) {
		h1 := elm.Find("h1")
		name = strings.Replace(h1.Text(), "\n", "", -1)
	})
	if name == "" {
		doc.Find("address").Each(func(arg1 int, elm *goquery.Selection) {
			meta, ok := elm.Find("meta").Attr("content")
			if ok {
				name = meta
			}
		})
	}
	return strings.TrimSpace(name)
}
