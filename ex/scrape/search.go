package scrape

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"gopkg.in/yaml.v2"

	"github.com/riyadennis/realty-tool/internal"
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
		return nil
	}
	cnf := &Config{}
	err = yaml.Unmarshal(file, cnf)
	if err != nil {
		log.Fatalf("invalid yaml :: %v", err)
		return nil
	}
	return cnf
}

// Search will run a search on the urls from config
func Search(urls []URL) []*internal.PropertyRecord {
	var records []*internal.PropertyRecord

	rDocument := propertyList(urls[1].Search)
	rIds := propertyIdsRightMove(rDocument)
	for _, id := range rIds {
		p := property(urls[1].View, id)
		if p != nil {
			records = append(records, p)
		}
	}

	wDocument := propertyList(urls[0].Search)
	wIds := propertyIdsWoodBury(wDocument)
	for _, id := range wIds {
		p := property(urls[0].View, id)
		if p != nil {
			records = append(records, p)
		}
	}
	return records
}

func propertyList(url string) *goquery.Document {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("error in accessing the site %v", err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		log.Fatalf("unable to perform thr searchs %v", resp.StatusCode)
	}
	document, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Fatalf("unable to read the content %v", err)
	}
	return document
}

func propertyIdsWoodBury(doc *goquery.Document) map[string]string {
	pids := make(map[string]string, 1)
	doc.Find("a").Each(func(index int, el *goquery.Selection) {
		href, exists := el.Attr("href")
		if exists {
			if index := strings.Index(href, "pid="); index > 0 {
				id := href[22:30]
				// check if id is already populated or not
				if ok := pids[id]; ok == "" {
					pids[id] = id
				}
			}
		}
	})
	return pids
}
func propertyIdsRightMove(doc *goquery.Document) map[string]string {
	pids := make(map[string]string, 1)
	doc.Find("div").Each(func(index int, el *goquery.Selection) {
		idStr, exists := el.Attr("id")
		if exists {
			index := strings.Index(idStr, "property-")
			if index >= 0 {
				id := idStr[9:17]
				if ok := pids[id]; ok == "" {
					pids[id] = fmt.Sprintf("%s.html", id)
				}
			}
		}
	})
	return pids
}
