package graph

import (
	"context"
	"errors"
	"log"
)

type Total struct {
	Count int `json:"count"`
}

func totalProperty(ctx context.Context, logger *log.Logger,  store *Store) (int, error){
	query := `query {
    	total: aggregateProperty{
      		count
    	}
	}`

	var Aggregate struct {
		Total `json:"total"`
	}

	if err := store.gql.Execute(ctx, query, &Aggregate); err != nil {
		logger.Printf("failed to execute graphql query in dgraph: %v", err)
		return 0, errors.New("no total found in response")
	}


	return Aggregate.Total.Count, nil
}



func totalArea(ctx context.Context, logger *log.Logger,  store *Store) (int, error){
	query := `query {
    	total: aggregateArea{
      		count
    	}
	}`

	var Aggregate struct {
		Total `json:"total"`
	}

	if err := store.gql.Execute(ctx, query, &Aggregate); err != nil {
		logger.Printf("failed to execute graphql query in dgraph: %v", err)
		return 0, errors.New("no total found in response")
	}


	return Aggregate.Total.Count, nil
}

func totalTransaction(ctx context.Context, logger *log.Logger,  store *Store) (int, error){
	query := `query {
    	total: aggregateTransaction{
      		count
    	}
	}`

	var Aggregate struct {
		Total `json:"total"`
	}

	if err := store.gql.Execute(ctx, query, &Aggregate); err != nil {
		logger.Printf("failed to execute graphql query in dgraph: %v", err)
		return 0, errors.New("no total found in response")
	}


	return Aggregate.Total.Count, nil
}


func totalPricePaid(ctx context.Context, logger *log.Logger,  store *Store) (int, error){
	query := `query {
    	total: aggregatePricePaidData{
      		count
    	}
	}`

	var Aggregate struct {
		Total `json:"total"`
	}

	if err := store.gql.Execute(ctx, query, &Aggregate); err != nil {
		logger.Printf("failed to execute graphql query in dgraph: %v", err)
		return 0, errors.New("no total found in response")
	}


	return Aggregate.Total.Count, nil
}
