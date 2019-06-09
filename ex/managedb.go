package ex

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/riyadennis/realty-tool/internal"
)

func checkAndUpdate(id, fullURL string, doc *goquery.Document) *internal.PropertyRecord {
	price := price(doc)
	status := status(doc)
	name := name(doc)
	p, err := internal.GetData(id)
	if err != nil {
		log.Fatalf("unable to fetch :: %v", err)
	}
	// if its a new property save
	if p == nil {
		p = &internal.PropertyRecord{
			ID:     id,
			Name:   name,
			Price:  price,
			Status: status,
			URL:    fullURL,
		}
		p.SavePayload()
		return p
	}
	p.Name = name
	p.Status = status
	if p.Price != price {
		p.Price = price
		err := internal.UpdatePrice(id, price)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	if p.Status != status {
		p.Status = status
		err := internal.UpdateStatus(id, status)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	return p
}
