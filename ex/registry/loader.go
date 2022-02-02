package registry

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"sync"
	"time"
)

type PricePaidData struct {
	ID           string
	dataSourceID string
	Property     *Property
}

type Property struct {
	ID              string
	Price           string
	TransactionDate time.Time
	Postcode        string
	Address         *Address
}

type Address struct {
	ID string
	// Primary Addressable Object Name (typically the house number or name)
	PAON string

	// Secondary Addressable Object Name – if there is a sub-building,
	//for example, the building is divided into flats, there will be a SAON
	SAON string

	Street string

	Locality string
	Town     string
	District string
	County   string
}

type Loader struct {
	PricePaidData sync.Map
	Property      sync.Map
	Address       sync.Map
}

func (l *Loader) LineToMutation(_ context.Context, errChan chan error, record []string) chan string {
	var (
		subMutationChan = make(chan string, len(record))
	)
	ppd, err := PPD(record)
	if err != nil {
		errChan <- err
	}
	if ppd == nil {
		return nil
	}

	subMutation := AddressMutation(ppd.Property.Address)
	subMutation += PropertyMutation(ppd.Property)
	subMutation += fmt.Sprintf(`
		_:%s <PricePaidData.property> _:%s  (dataSourceID=%q) .
		_:%s <PricePaidData.dataSourceID> %q .
		_:%s <dgraph.type> "PricePaidData" .
	`,
		ppd.ID, ppd.Property.ID, ppd.dataSourceID,
		ppd.ID, ppd.dataSourceID,
		ppd.ID)

	subMutationChan <- subMutation

	return subMutationChan
}

func PropertyMutation(property *Property) string {
	return fmt.Sprintf(`
		_:%s <Property.address> _:%s .
		_:%s <Property.price> %q .
		_:%s <Property.transactionDate> %q .
		_:%s <Property.postcode> %q .
		_:%s <dgraph.type> "Property" .
`,
		property.ID, property.Address.ID,
		property.ID, property.Price,
		property.ID, property.TransactionDate,
		property.ID, property.Postcode,
		property.ID)
}

func AddressMutation(addr *Address) string {
	return fmt.Sprintf(`	
		_:%s <Address.paon> %q .
		_:%s <Address.saon> %q .
		_:%s <Address.street> %q .
		_:%s <Address.locality> %q .
		_:%s <Address.town> %q .
		_:%s <Address.district> %q .
		_:%s <Address.county> %q .
		_:%s <dgraph.type> "Address" .`,
		addr.ID, addr.PAON,
		addr.ID, addr.SAON,
		addr.ID, addr.Street,
		addr.ID, addr.Locality,
		addr.ID, addr.Town,
		addr.ID, addr.District,
		addr.ID, addr.County,
		addr.ID)
}

func PPD(record []string) (*PricePaidData, error) {
	ppdID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}

	transactionDate, err := time.Parse("2006-01-02 15:04", record[2])
	if err != nil {
		return nil, err
	}

	return &PricePaidData{
		ID:           "ppd" + ppdID.String(),
		dataSourceID: record[0],
		Property: &Property{
			ID:              "p" + ppdID.String(),
			Price:           record[1],
			TransactionDate: transactionDate,
			Postcode:        record[3],
			Address: &Address{
				ID:       "addr" + ppdID.String(),
				PAON:     record[4],
				SAON:     record[5],
				Street:   record[6],
				Locality: record[7],
				Town:     record[8],
				District: record[9],
				County:   record[10],
			},
		},
	}, nil

}
