package ex

import (
	"log"

	"github.com/PuerkitoBio/goquery"
	"github.com/riyadennis/realty-tool/internal"
)

func checkAndUpdate(id, fullURL string, doc *goquery.Document) {
	price := price(doc)
	status := status(doc)
	p, err := internal.GetData(id)
	if err != nil {
		log.Fatalf("unable to fetch :: %v", err)
	}
	// if its a new property save
	if p == nil {
		pa := &internal.PropertyRecord{
			ID:     id,
			Name:   name(doc),
			Price:  price,
			Status: status,
			URL:    fullURL,
		}
		pa.SavePayload()
		return
	}
	if p.Price != price {
		err := internal.UpdatePrice(id, price)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	if p.Status != status {
		err := internal.UpdateStatus(id, status)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	return
}
