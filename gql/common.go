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

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) _(ctx context.Context) (*bool, error) {
	panic(fmt.Errorf("not implemented: _ - _"))
}
*/
