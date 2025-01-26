package gql

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"fmt"
	"log"

	"github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/laclipasa/la-clipasa/ent"
	"github.com/laclipasa/la-clipasa/ent/user"
	"github.com/laclipasa/la-clipasa/internal"
	"github.com/laclipasa/la-clipasa/internal/auth"
)

type Resolver struct {
	ent *ent.Client
}

func NewResolver(entClient *ent.Client) Config {
	return Config{
		Resolvers: &Resolver{
			ent: entClient,
		},
		Directives: DirectiveRoot{
			HasRole: func(ctx context.Context, obj any, next graphql.Resolver, role user.Role) (res any, err error) {
				log.Default().Printf("required role: %s", role)
				u, ok := internal.GetUserFromContext(ctx)
				if !ok {
					return nil, fmt.Errorf("unauthenticated")
				}

				if auth.RoleRank.Get(u.Role) < auth.RoleRank.Get(role) {
					return nil, fmt.Errorf("unauthorized")
				}

				return next(ctx)
			},
		},
	}
}

func NewServer(ctx context.Context) *handler.Server {
	resolver := NewResolver(ent.FromContext(ctx))
	return handler.NewDefaultServer(NewExecutableSchema(resolver))
}
