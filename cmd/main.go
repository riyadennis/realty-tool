package main

import (
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/realty-tool/ex/registry"
)

func main() {
	//urls := scrape.GetURLS("config.yaml")
	//if urls != nil {
	//	records := scrape.Search(urls.Urls)
	//	if len(records) == 0 {
	//		panic("no records")
	//	}
	//	err := scrape.CheckAndUpdate(records)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	logger := log.New(os.Stdout, "LOADER: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	loader := &registry.Loader{}
	err := registry.LoadFromCSV(logger, "data/input.csv", "output.rdf", loader)
	if err != nil {
		panic(err)
	}
}
