package gql

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/laclipasa/la-clipasa/ent"
	"github.com/laclipasa/la-clipasa/gql/model"
)

// CreateUser is the resolver for the createUser field.
func (r *mutationResolver) CreateUser(ctx context.Context, input ent.CreateUserInput) (*ent.User, error) {
	return r.ent.User.Create().SetInput(input).Save(ctx)
}

// UpdateUser is the resolver for the updateUser field.
func (r *mutationResolver) UpdateUser(ctx context.Context, id int, input ent.UpdateUserInput) (*ent.User, error) {
	panic(fmt.Errorf("not implemented: UpdateUser - updateUser"))
}

// AdminUpdateUser is the resolver for the adminUpdateUser field.
func (r *mutationResolver) AdminUpdateUser(ctx context.Context, id int, input model.AdminUpdateUserInput) (*ent.User, error) {
	return r.ent.User.UpdateOneID(id).SetInput(*input.BaseInput).Save(ctx)
}

// DeleteUser is the resolver for the deleteUser field.
func (r *mutationResolver) DeleteUser(ctx context.Context, id int) (bool, error) {
	panic(fmt.Errorf("not implemented: DeleteUser - deleteUser"))
}

// PaginatedUsers is the resolver for the paginatedUsers field.
func (r *queryResolver) PaginatedUsers(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int) (*ent.UserConnection, error) {
	panic(fmt.Errorf("not implemented: PaginatedUsers - paginatedUsers"))
}
