package ex

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"
)

// Config contains the url strings which we need to search
type Config struct {
	Urls []URL `yaml:"urls"`
}

// URL holds the search and view page
type URL struct {
	Search string `yaml:"search"`
	View   string `yaml:"view"`
}

//GetURLS gets all the urls that is in the config file to search
func GetURLS(path string) *Config {
	file, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("unable to open config file :: %v", err)
	}
	cnf := &Config{}
	err = yaml.Unmarshal(file, cnf)
	if err != nil {
		log.Fatalf("invalid yaml :: %v", err)
	}
	return cnf
}

// Search will run a search on the urls from config
func Search(urls []URL) {
	for _, u := range urls {
		document := propertyList(u.Search)
		ids := propertyIds(document)
		for _, id := range ids {
			viewProperty(u.View, id)
			fmt.Println("*********************")
		}
	}
}

func propertyList(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error in accessing the site %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("invalid status code %v", resp.StatusCode)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("unable to read the content %v", err)
	}
	return document
}

func propertyIds(doc *goquery.Document) map[string]string {
	pids := make(map[string]string, 1)
	doc.Find("a").Each(func(index int, el *goquery.Selection) {
		href, exists := el.Attr("href")
		if exists {
			if index := strings.Index(href, "pid="); index > 0 {
				id := href[22 : 22+8]
				// check if id is already populated or not
				if ok := pids[id]; ok == "" {
					pids[id] = id
				}
			}
		}
	})
	return pids
}
