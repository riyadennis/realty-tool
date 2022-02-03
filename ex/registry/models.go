package registry

import "time"

type PricePaidData struct {
	ID           string
	DataSourceID string
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

