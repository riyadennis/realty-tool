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
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		log.Fatalf("unable to view property :: %v", err)
	}
	re := regexp.MustCompile("[0-9]+")
	p := propertyRecord{
		id:     re.FindAllString(id, -1)[0],
		url:    fullURL,
		price:  price(doc),
		status: status(doc),
		name:   name(doc),
	}
	fmt.Printf("url : %s, price : %v\n", fullURL, p)
	return p.checkAndUpdate()
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
