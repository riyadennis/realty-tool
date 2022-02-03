package registry

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"log"
	"strings"
	"sync"
	"time"
)


type PricePaidData struct {
	ID           string
	dataSourceID string
	Transactions []*Transaction
}

type Transaction struct {
	TransactionDate time.Time
	Price string
	Property     *Property
}

type Property struct {
	ID              string
	DoorNumber 		string
	Neighbourhood *Neighbourhood
	Area         	*Area
	Postcode        string
	PricePaidData []*PricePaidData
}

type Area struct {
	ID string
	Locality string
	Town     string
	District string
	County   string
	OutCode string
}

type Neighbourhood struct {
	Street string
	Postcodes []string
}

type Loader struct {
	Logger *log.Logger
	PricePaidData sync.Map
	Property      sync.Map
	Area       sync.Map
}

func (l *Loader) LineToMutation(_ context.Context, _ chan error, record []string) chan string {
	var (
		subMutationChan = make(chan string, len(record))
		ppd *PricePaidData
		area *Area
		subMutation string

	)
	go func(){

		property, err := PropertyFromCSV(record)
		if err != nil {
			l.Logger.Fatalf("error reading property from CSV : %v", err)
		}

		data, ok := l.PricePaidData.LoadOrStore(record[0], ppd)
		if ok && data != nil {
			ppd, _ = data.(*PricePaidData)
		} else {
			ppdID, err := uuid.NewUUID()
			if err != nil {
				l.Logger.Fatalf("error creating uuid for ppd : %v", err)
			}
			ppd = &PricePaidData{
				ID: ppdID.String(),
				dataSourceID: record[0],
			}

			subMutation += fmt.Sprintf(`
		_:%s <PricePaidData.DataSourceID> %q .
		_:%s <dgraph.type> "PricePaidData" .
	`,
				ppd.ID,ppd.dataSourceID,
				ppd.ID)
		}

		transactionDate, err := time.Parse("2006-01-02 15:04", record[2])
		if err != nil {
			l.Logger.Fatalf("error creating transaction date : %v", err)
		}

		transaction :=  &Transaction{
			TransactionDate: transactionDate,
			Price: record[1],
			Property: property,
		}

		//ppd.Transactions = append(ppd.Transactions,transaction)
		subMutation += fmt.Sprintf(`
		_:%s <Transaction.Property> _:%s  (dataSourceID=%q) .
		_:%s <Transactions.TransactionDate> %q .
		_:%s <Transaction.Price> %q .
		_:%s <dgraph.type> "Transaction" .
		_:%s <PricePaidData.Transactions> _:%s .
`,
			"trans"+property.ID, property.ID, ppd.dataSourceID,
			"trans"+property.ID, transaction.TransactionDate,
			"trans"+property.ID, transaction.Price,
			"trans"+property.ID,
			ppd.ID, "trans"+property.ID,
		)

		dataArea, ok := l.Area.LoadOrStore(property.Area.OutCode, property.Area)
		if ok && dataArea != nil{
			area, _ = dataArea.(*Area)
		}else{
			area = property.Area
			subMutation += AreaMutation(area)
		}

		dataProperty, ok := l.Property.LoadOrStore(property.Postcode+property.DoorNumber+property.Neighbourhood.Street, property)
		if ok && dataProperty != nil{
			property, _ = dataProperty.(*Property)
		}else{
			subMutation += PropertyMutation(property, property.ID)
		}

		subMutation += fmt.Sprintf(`
	_:%s <Property.Area> _:%s .
	_:%s <Property.PricePaidData> _:%s .
	`,
			property.ID, area.ID,
			property.ID, ppd.ID)

		subMutationChan <- subMutation
	}()

	return subMutationChan
}

func PropertyMutation(property *Property, ppdID string) string {
	return fmt.Sprintf(`
		_:%s <Property.DoorNumber> %q .
		_:%s <Property.Postcode> %q .
		_:%s <dgraph.type> "Property" .
`,
		property.ID, property.DoorNumber,
		property.ID, property.Postcode,
		property.ID)
}

func AreaMutation(area *Area) string {
	return fmt.Sprintf(`
		_:%s <Area.Locality> %q .
		_:%s <Area.Town> %q .
		_:%s <Area.District> %q .
		_:%s <Area.County> %q .
		_:%s <Area.OutCode> %q .
		_:%s <dgraph.type> "Area" .`,
		area.ID, area.Locality,
		area.ID, area.Town,
		area.ID, area.District,
		area.ID, area.County,
		area.ID, area.OutCode,
		area.ID)
}


func PropertyFromCSV( record []string) (*Property, error) {
	var (
		postcode string
		outcode string
	)
	ppdID, err := uuid.NewUUID()
	if err != nil {
		return nil, err
	}


	postcode = record[3]

	if postcode != ""{
		outCodeArr := strings.Split(postcode," ")
		if (len(outCodeArr)) >0{
			outcode = outCodeArr[0]
		}
	}

	return &Property{
		ID:               ppdID.String(),
		DoorNumber:       record[7],
		Postcode:        postcode,
		Area: &Area{
			ID:       "addr" + ppdID.String(),
			Locality: record[10],
			Town:     record[11],
			District: record[12],
			County:   record[13],
			OutCode: outcode,
		},
		Neighbourhood: &Neighbourhood{
			Postcodes: []string{record[3]},
			Street:   record[9],
		},
	}, nil


}
