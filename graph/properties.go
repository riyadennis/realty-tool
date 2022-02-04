package graph

import (
	"context"
	"errors"
	"fmt"
	"github.com/riyadennis/realty-tool/graph/model"
	"log"
)

func Properties(ctx context.Context, logger *log.Logger,  store *Store, postcode string) ([]*model.Property, error){
	query :=  fmt.Sprintf(`
	query{
		properties: queryProperty(filter: {Postcode: {alloftext : %q}}){
			DoorNumber
			Postcode
			id
			Area{
				Locality
				Town
				County
			}
			PricePaidData{
				DataSourceID
				Transactions{
					TransactionDate
					Price
				}
			}
		}	
	}`, postcode)

	var Result struct {
		Properties []*model.Property `json:"properties"`
	}

	if err := store.gql.Execute(ctx, query, &Result); err != nil {
		logger.Printf("failed to execute graphql query in dgraph: %v", err)
		return nil, errors.New("no total found in response")
	}

	return Result.Properties, nil
}
