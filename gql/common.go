package gql

import (
	"context"
	"fmt"
)

// M is the resolver for the _m field.
func (r *mutationResolver) M(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: M - _m"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
