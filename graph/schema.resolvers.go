package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/riyadennis/realty-tool/graph/generated"
	"github.com/riyadennis/realty-tool/graph/model"
)

func (r *queryResolver) TotalProperties(ctx context.Context) (int, error) {
	return totalProperty(ctx,r.Logger,  r.Store)
}

func (r *queryResolver) TotalArea(ctx context.Context) (int, error) {
	return totalArea(ctx, r.Logger, r.Store)
}

func (r *queryResolver) TotalTransaction(ctx context.Context) (int, error) {
	return totalTransaction(ctx, r.Logger, r.Store)
}

func (r *queryResolver) TotalPricePaidData(ctx context.Context) (int, error) {
	return totalPricePaid(ctx, r.Logger, r.Store)
}

func (r *queryResolver) Properties(ctx context.Context) ([]*model.Property, error) {
	return []*model.Property{
		{
			DoorNumber: "test",
		},
	}, nil
}

func (r *transactionResolver) TransactionDate(ctx context.Context, obj *model.Transaction) (string, error) {
	panic(fmt.Errorf("not implemented"))
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type queryResolver struct{ *Resolver }
type transactionResolver struct{ *Resolver }
