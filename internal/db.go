package internal

import (
	"database/sql"
	"fmt"
	"time"

	"github.com/golang-migrate/migrate"
	"github.com/golang-migrate/migrate/database/mysql"
	"github.com/sirupsen/logrus"
)
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

func setUpForMigration(db *sql.DB)(*migrate.Migrate){
	migrationConfig := &mysql.Config{}
	driver, _ := mysql.WithInstance(db, migrationConfig)
	migrate, err := migrate.NewWithDatabaseInstance(
		"file://internal/migrations/",
		"properties",
		driver,
	)
	if err != nil {
		logrus.Fatalf("%v", err)
	}
	return migrate
}

// MigrateUp creates the needed tables for the first time
func MigrateUp() (bool) {
	db, err := InitDB()
	if err != nil {
		return false
	}
	migrate:=setUpForMigration(db)
	err = migrate.Steps(1)
	if err != nil{
		logrus.Fatalf("failed migration :: %v", err)
	}
	migrate.Up()
	fmt.Println("tables created")
	return true
}

// MigrateDown removes tables will loose all the data if this is done
func MigrateDown() (bool){
	db, err := InitDB()
	if err != nil {
		return false
	}
	fmt.Println("Undoing  migrations")
	migrate := setUpForMigration(db)
	err = migrate.Steps(1)
	if err != nil{
		logrus.Fatalf("failed migration :: %v", err)
	}
	migrate.Down()
	fmt.Println("Done")
	return true
}

func getCurrentTimeStamp() string {
	return time.Now().Format("2006-01-02 15:04:05")
}
