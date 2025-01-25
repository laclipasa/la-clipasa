package gql

import (
	"context"
	"fmt"
)

// Empty is the resolver for the _empty field.
func (r *mutationResolver) Empty(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Empty - _empty"))
}

// Empty is the resolver for the _empty field.
func (r *queryResolver) Empty(ctx context.Context) (*string, error) {
	panic(fmt.Errorf("not implemented: Empty - _empty"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
