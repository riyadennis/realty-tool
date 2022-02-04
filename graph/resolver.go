package graph

import "log"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{
	Store *Store
	Logger *log.Logger
}

func NewResolver(store *Store, logger *log.Logger) *Resolver{
	return &Resolver{
		Store: store,
		Logger: logger,
	}
}
