package gql

import (
	"context"
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/laclipasa/la-clipasa/ent"
	"github.com/laclipasa/la-clipasa/ent/note"
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

// NodeID is the resolver for the nodeId field.
func (r *noteResolver) NodeID(ctx context.Context, obj *ent.Note) (string, error) {
	return base64.RawURLEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%d", note.Table, obj.ID))), nil
}

// BodyMarkdown is the resolver for the bodyMarkdown field.
func (r *noteResolver) BodyMarkdown(ctx context.Context, obj *ent.Note) (string, error) {
	return obj.Body, nil
}

// BodyHTML is the resolver for the bodyHtml field.
func (r *noteResolver) BodyHTML(ctx context.Context, obj *ent.Note) (string, error) {
	panic(fmt.Errorf("not implemented: BodyHTML - bodyHtml"))
}

// Notes is the resolver for the notes field.
func (r *queryResolver) Notes(ctx context.Context) ([]*ent.Note, error) {
	return r.ent.Note.Query().All(ctx)
}

// Node is the resolver for the node field.
// Node is the resolver for the node field.
func (r *queryResolver) Node(ctx context.Context, nodeID string) (ent.Noder, error) {
	rawNodeID, err := base64.RawURLEncoding.DecodeString(nodeID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse node id: base64 decode error")
	}

	splitNodeID := strings.Split(string(rawNodeID), ":")
	if len(splitNodeID) != 2 {
		return nil, fmt.Errorf("failed to parse node id: wrong number of parts")
	}

	switch splitNodeID[0] {
	// add other cases here (custom table names, non-ent types, etc.)

	case note.Table:
		id, err := strconv.Atoi(splitNodeID[1])
		if err != nil {
			return nil, err
		}
		return r.ent.Noder(ctx, id, ent.WithFixedNodeType(splitNodeID[0]))

	default:
		return nil, fmt.Errorf("failed parse node id type")
	}
}

// Note returns NoteResolver implementation.
func (r *Resolver) Note() NoteResolver { return &noteResolver{r} }

type noteResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
/*
	func (r *mutationResolver) Node(ctx context.Context, nodeID string) (ent.Noder, error) {
	panic(fmt.Errorf("not implemented: Node - node"))
}
*/
