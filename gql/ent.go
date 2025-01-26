package gql

import (
	"context"
	"fmt"

	"entgo.io/contrib/entgql"
	"github.com/laclipasa/la-clipasa/ent"
)

// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, id string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}

// Nodes is the resolver for the nodes field.
func (r *queryResolver) Nodes(ctx context.Context, ids []string) ([]ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Nodes - nodes"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.CommentWhereInput) (*ent.CommentConnection, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, orderBy *ent.NoteOrder, where *ent.NoteWhereInput) (*ent.NoteConnection, error) {
	panic(fmt.Errorf("not implemented: Notes - notes"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.PostWhereInput) (*ent.PostConnection, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context, after *entgql.Cursor[int], first *int, before *entgql.Cursor[int], last *int, where *ent.UserWhereInput) (*ent.UserConnection, error) {
	return r.ent.User.Query().Paginate(ctx, after, first, before, last)
}

// Note returns NoteResolver implementation.
func (r *Resolver) Note() NoteResolver { return &noteResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type noteResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
