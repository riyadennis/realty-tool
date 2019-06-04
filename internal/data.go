package internal

import (
	"github.com/sirupsen/logrus"
)

// PropertyRecord holds data for individual property
type PropertyRecord struct {
	ID       string
	Name     string
	Price    string
	URL      string
	Status   string
	NewPrice string
}

// SavePayload saves information about individual properties
func (pr *PropertyRecord) SavePayload() error {
	db, err := InitDB()
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	query, err := db.Prepare("INSERT INTO property_data(id,price,name,url,status,InsertedDatetime) VALUES(?, ?, ?, ?,?,?)")
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	_, err = query.Exec(pr.ID, pr.Price, pr.Name, pr.URL, pr.Status, getCurrentTimeStamp())
	if err != nil {
		logrus.Errorf("Unable to save payload %s", err.Error())
		return err
	}
	return nil
}

// UpdatePrice saves information about individual properties
func UpdatePrice(id, price string) error {
	db, err := InitDB()
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	query, err := db.Prepare("UPDATE property_data SET new_price = ?, price_updated= ? WHERE id = ?")
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	_, err = query.Exec(price, getCurrentTimeStamp(), id)
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	return nil
}

// UpdateStatus saves information about individual properties
func UpdateStatus(id, status string) error {
	db, err := InitDB()
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	query, err := db.Prepare("UPDATE property_data SET status = ?,status_updated= ? WHERE id = ?")
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	_, err = query.Exec(status, getCurrentTimeStamp(), id)
	if err != nil {
		logrus.Errorf("Unable to update payload %s", err.Error())
		return err
	}
	return nil
}

func GetData(id string) (*PropertyRecord, error) {
	var price, status string
	db, err := InitDB()
	defer db.Close()
	if err != nil {
		logrus.Errorf("Unable to get data from db %s", err.Error())
		return nil, err
	}
	query := "SELECT price, status from property_data where id = '" + id + "'"
	row := db.QueryRow(query)
	err = row.Scan(&price, &status)
	if err != nil {
		logrus.Errorf("Unable to get  details %s", err.Error())
		return nil, nil
	}
	if price != "" {
		return &PropertyRecord{Price: price, Status: status}, nil
	}
	return nil, nil
}
