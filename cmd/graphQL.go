package main

import (
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/golang-migrate/migrate/source/file"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/riyadennis/realty-tool/graph"
	"github.com/riyadennis/realty-tool/graph/generated"
)

const (
	defaultPort = "8081"
)

func main() {
	logger := log.New(os.Stdout, "graphQL: ", log.LstdFlags|log.Lmicroseconds|log.Lshortfile)
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
