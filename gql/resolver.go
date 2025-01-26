package gql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/laclipasa/la-clipasa/ent"
)

type Resolver struct {
	ent *ent.Client
}

func NewResolver(entClient *ent.Client) Config {
	return Config{
		Resolvers: &Resolver{
			ent: entClient,
		},
	}
}

func NewServer(ctx context.Context) *handler.Server {
	resolver := NewResolver(ent.FromContext(ctx))
	return handler.NewDefaultServer(NewExecutableSchema(resolver))
}
