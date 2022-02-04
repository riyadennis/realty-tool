package main

import (
	"log"
	"net/http"
	"os"
	"sync"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/riyadennis/realty-tool/ex/registry"
	"github.com/riyadennis/realty-tool/graph"
	"github.com/riyadennis/realty-tool/graph/generated"
)

const (
	defaultPort = "8081"
)

func main() {
	logger := log.New(os.Stdout, "LOADER: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)

	if os.Getenv("mode") == "load"{
		loader := &registry.Loader{
			Logger: logger,
			PricePaidData: sync.Map{},
			Property: sync.Map{},
			Area: sync.Map{},
		}
		err := registry.LoadFromCSV(logger, "data/pp-2020.csv", "output.rdf", loader)
		if err != nil {
			logger.Fatalf("error loading CSV file to rdf: %v", err)
		}
	}



	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	store := graph.NewStore(logger, graph.NewGraphQL(&graph.GraphQLConfig{
		URL: "http://localhost:8080",
	}))

	resolver := graph.NewResolver(store, logger)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
