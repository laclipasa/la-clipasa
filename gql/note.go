package gql

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/laclipasa/la-clipasa/ent"
	"github.com/laclipasa/la-clipasa/gql/model"
)

// CreateNote is the resolver for the createNote field.
func (r *mutationResolver) CreateNote(ctx context.Context, input model.NoteInput) (*ent.Note, error) {
	return r.ent.Note.Create().
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
}

// UpdateNote is the resolver for the updateNote field.
func (r *mutationResolver) UpdateNote(ctx context.Context, id int, input model.NoteInput) (*ent.Note, error) {
	return r.ent.Note.UpdateOneID(id).
		SetTitle(input.Title).
		SetBody(input.Body).
		Save(ctx)
}

// DeleteNote is the resolver for the deleteNote field.
func (r *mutationResolver) DeleteNote(ctx context.Context, id int) (bool, error) {
	err := r.ent.Note.DeleteOneID(id).Exec(ctx)
	if err != nil {
		if ent.IsNotFound(err) {
			return false, nil
		}

		return false, err
	}

	return true, nil
}

// PostReadHours is the resolver for the postReadHours field.
func (r *noteResolver) PostReadHours(ctx context.Context, obj *ent.Note) (int, error) {
	panic(fmt.Errorf("not implemented: PostReadHours - postReadHours"))
}

// PaginatedNotes is the resolver for the paginatedNotes field.
func (r *queryResolver) PaginatedNotes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int) (*ent.NoteConnection, error) {
	return r.ent.Note.Query().Paginate(ctx, after, first, before, last)
}
