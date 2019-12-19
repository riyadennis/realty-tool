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

func CheckAndUpdate(record []*internal.PropertyRecord) error {
	err := internal.Init()
	if err != nil {
		log.Fatalf("unable to  initialise :: %v", err)
		return err
	}
	for _, pr := range record {
		p, err := internal.Property(pr.ID)
		if err != nil {
			log.Fatalf("unable to fetch :: %v", err)
			continue
		}
		// if its a new property save
		if p == nil {
			pr.SavePayload()
		}
		if pr == nil {
			continue
		}
		if p.Price != pr.Price {
			p.Price = pr.Price
			err := internal.UpdatePrice(pr.ID, pr.Price)
			if err != nil {
				log.Fatalf("unable to update price :: %v", err)
				return err
			}
		}
		if p.Status != pr.Status {
			p.Status = pr.Status
			err := internal.UpdateStatus(pr.ID, pr.Status)
			if err != nil {
				log.Fatalf("unable to update price :: %v", err)
				return err
			}
		}
	}
	return nil
}
