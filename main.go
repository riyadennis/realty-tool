package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/realty-tool/ex"
)

func main() {
	// migrateFlag := flag.String("setup", "up", "To Create tables if they do not exist ")
	// flag.Parse()
	// if *migrateFlag == "up" {
	// 	created := internal.MigrateUp()
	// 	if !created {
	// 		logrus.Fatal("unable to create the tables")
	// 	}
	// }
	// if *migrateFlag == "down" {
	// 	created := internal.MigrateDown()
	// 	if !created {
	// 		logrus.Fatal("unable to create the tables")
	// 	}
	// }
	urls := ex.GetURLS("config.yaml")
	ex.Search(urls.Urls)
}
