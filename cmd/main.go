package main

import (
	"log"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"github.com/riyadennis/realty-tool/ex/registry"
)

func main() {
	logger := log.New(os.Stdout, "LOADER: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	loader := &registry.Loader{
		Logger: logger,
		PricePaidData: sync.Map{},
		Property: sync.Map{},
		Area: sync.Map{},
	}
	err := registry.LoadFromCSV(logger, "data/pp-2020.csv", "output.rdf", loader)
	if err != nil {
		panic(err)
	}
}
