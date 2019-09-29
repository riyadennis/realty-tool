package ex

import (
	"log"

	"github.com/riyadennis/realty-tool/internal"
)

type propertyRecord struct {
	id     string
	url    string
	price  string
	status string
	name   string
}

func (pr *propertyRecord) checkAndUpdate() *internal.PropertyRecord {
	err := internal.Init()
	if err != nil {
		panic(err)
	}
	p, err := internal.GetData(pr.id)
	if err != nil {
		log.Fatalf("unable to fetch :: %v", err)
	}
	// if its a new property save
	if p == nil {
		p = &internal.PropertyRecord{
			ID:     pr.id,
			Name:   pr.name,
			Price:  pr.price,
			Status: pr.status,
			URL:    pr.url,
		}
		p.SavePayload()
		return p
	}
	p.Name = pr.name
	p.Status = pr.status
	if p.Price != pr.price {
		p.Price = pr.price
		err := internal.UpdatePrice(pr.id, pr.price)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	if p.Status != pr.status {
		p.Status = pr.status
		err := internal.UpdateStatus(pr.id, pr.status)
		if err != nil {
			log.Fatalf("unable to update price :: %v", err)
		}
	}
	return p
}
