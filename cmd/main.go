package main

import (
	"log"
	"os"

	"github.com/riyadennis/realty-tool/ex/scrape"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/realty-tool/ex/registry"
)

func main() {
	logger := log.New(os.Stdout, "LOADER: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	urls := scrape.GetURLS(logger, "config.yaml")
	if urls != nil {
		records := scrape.Search(logger, urls.Urls)
		if len(records) == 0 {
			panic("no records")
		}
		err := scrape.CheckAndUpdate(records)
		if err != nil {
			panic(err)
		}
	}

	loader := &registry.Loader{}
	err := registry.LoadFromCSV(logger, "data/input.csv", "output.rdf", loader)
	if err != nil {
		panic(err)
	}
}
