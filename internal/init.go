package internal

import (
	"database/sql"
	"fmt"

	"github.com/sirupsen/logrus"
)

var db *sql.DB

// Init does all the initialisation
func Init() error {
	var err error
	db, err = InitDB()
	if err != nil {
		return err
	}
	return nil
}

// InitDB initailises the db connection
func InitDB() (*sql.DB, error) {
	dbConnectionString := fmt.Sprintf("%s:%s@/%s?multiStatements=true", "root", "", "properties")
	dbConnector, err := sql.Open("mysql", dbConnectionString)
	if err != nil {
		logrus.Errorf("Unable to start database %s", err.Error())
		return nil, err
	}
	return dbConnector, nil
}
