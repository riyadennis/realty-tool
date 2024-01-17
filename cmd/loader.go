package main

import (
	"log"
	"os"
	"sync"

	"github.com/riyadennis/realty-tool/ex/registry"
)

func main() {
	logger := log.New(os.Stdout, "Loader: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
	loader := &registry.Loader{
		Logger:        logger,
		PricePaidData: sync.Map{},
		Property:      sync.Map{},
		Area:          sync.Map{},
	}
	err := registry.LoadFromCSV(logger, "data/pp-2020.csv", "output.rdf", loader)
	if err != nil {
		logger.Fatalf("error loading CSV file to rdf: %v", err)
	}
}
