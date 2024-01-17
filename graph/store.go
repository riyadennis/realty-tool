package graph

import (
	"log"
	"net"
	"net/http"
	"time"

	"github.com/ardanlabs/graphql"
)

// Store manages the set of API's for city access.
type Store struct {
	log *log.Logger
	gql *graphql.GraphQL
}

// GraphQLConfig represents comfiguration needed to support managing, mutating,
// and querying the database.
type GraphQLConfig struct {
	URL             string
	AuthHeaderName  string
	AuthToken       string
	CloudHeaderName string
	CloudToken      string
}

// NewStore constructs a weather store for api access.
func NewStore(log *log.Logger, gql *graphql.GraphQL) *Store {
	return &Store{
		log: log,
		gql: gql,
	}
}

// NewGraphQL constructs a graphql value for use to access the databse.
func NewGraphQL(gqlConfig *GraphQLConfig) *graphql.GraphQL {
	client := http.Client{
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   30 * time.Second,
				KeepAlive: 30 * time.Second,
			}).DialContext,
			ForceAttemptHTTP2:     true,
			MaxIdleConns:          100,
			IdleConnTimeout:       90 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 1 * time.Second,
		},
	}

	graphql := graphql.New(gqlConfig.URL,
		graphql.WithClient(&client),
		graphql.WithHeader(gqlConfig.AuthHeaderName, gqlConfig.AuthToken),
		graphql.WithHeader(gqlConfig.CloudHeaderName, gqlConfig.CloudToken),
	)

	return graphql
}
