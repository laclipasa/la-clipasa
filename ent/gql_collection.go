// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/contrib/entgql"
	"github.com/99designs/gqlgen/graphql"
	"github.com/laclipasa/la-clipasa/ent/note"
)

// CollectFields tells the query-builder to eagerly load connected nodes by resolver context.
func (n *NoteQuery) CollectFields(ctx context.Context, satisfies ...string) (*NoteQuery, error) {
	fc := graphql.GetFieldContext(ctx)
	if fc == nil {
		return n, nil
	}
	if err := n.collectField(ctx, false, graphql.GetOperationContext(ctx), fc.Field, nil, satisfies...); err != nil {
		return nil, err
	}
	return n, nil
}

func (n *NoteQuery) collectField(ctx context.Context, oneNode bool, opCtx *graphql.OperationContext, collected graphql.CollectedField, path []string, satisfies ...string) error {
	path = append([]string(nil), path...)
	var (
		unknownSeen    bool
		fieldSeen      = make(map[string]struct{}, len(note.Columns))
		selectedFields = []string{note.FieldID}
	)
	for _, field := range graphql.CollectFields(opCtx, collected.Selections, satisfies) {
		switch field.Name {
		case "title":
			if _, ok := fieldSeen[note.FieldTitle]; !ok {
				selectedFields = append(selectedFields, note.FieldTitle)
				fieldSeen[note.FieldTitle] = struct{}{}
			}
		case "body":
			if _, ok := fieldSeen[note.FieldBody]; !ok {
				selectedFields = append(selectedFields, note.FieldBody)
				fieldSeen[note.FieldBody] = struct{}{}
			}
		case "createdat":
			if _, ok := fieldSeen[note.FieldCreatedAt]; !ok {
				selectedFields = append(selectedFields, note.FieldCreatedAt)
				fieldSeen[note.FieldCreatedAt] = struct{}{}
			}
		case "updatedat":
			if _, ok := fieldSeen[note.FieldUpdatedAt]; !ok {
				selectedFields = append(selectedFields, note.FieldUpdatedAt)
				fieldSeen[note.FieldUpdatedAt] = struct{}{}
			}
		case "id":
		case "__typename":
		default:
			unknownSeen = true
		}
	}
	if !unknownSeen {
		n.Select(selectedFields...)
	}
	return nil
}

type notePaginateArgs struct {
	first, last   *int
	after, before *Cursor
	opts          []NotePaginateOption
}

func newNotePaginateArgs(rv map[string]any) *notePaginateArgs {
	args := &notePaginateArgs{}
	if rv == nil {
		return args
	}
	if v := rv[firstField]; v != nil {
		args.first = v.(*int)
	}
	if v := rv[lastField]; v != nil {
		args.last = v.(*int)
	}
	if v := rv[afterField]; v != nil {
		args.after = v.(*Cursor)
	}
	if v := rv[beforeField]; v != nil {
		args.before = v.(*Cursor)
	}
	if v, ok := rv[orderByField]; ok {
		switch v := v.(type) {
		case map[string]any:
			var (
				err1, err2 error
				order      = &NoteOrder{Field: &NoteOrderField{}, Direction: entgql.OrderDirectionAsc}
			)
			if d, ok := v[directionField]; ok {
				err1 = order.Direction.UnmarshalGQL(d)
			}
			if f, ok := v[fieldField]; ok {
				err2 = order.Field.UnmarshalGQL(f)
			}
			if err1 == nil && err2 == nil {
				args.opts = append(args.opts, WithNoteOrder(order))
			}
		case *NoteOrder:
			if v != nil {
				args.opts = append(args.opts, WithNoteOrder(v))
			}
		}
	}
	if v, ok := rv[whereField].(*NoteWhereInput); ok {
		args.opts = append(args.opts, WithNoteFilter(v.Filter))
	}
	return args
}

const (
	afterField     = "after"
	firstField     = "first"
	beforeField    = "before"
	lastField      = "last"
	orderByField   = "orderBy"
	directionField = "direction"
	fieldField     = "field"
	whereField     = "where"
)

func fieldArgs(ctx context.Context, whereInput any, path ...string) map[string]any {
	field := collectedField(ctx, path...)
	if field == nil || field.Arguments == nil {
		return nil
	}
	oc := graphql.GetOperationContext(ctx)
	args := field.ArgumentMap(oc.Variables)
	return unmarshalArgs(ctx, whereInput, args)
}

// unmarshalArgs allows extracting the field arguments from their raw representation.
func unmarshalArgs(ctx context.Context, whereInput any, args map[string]any) map[string]any {
	for _, k := range []string{firstField, lastField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		i, err := graphql.UnmarshalInt(v)
		if err == nil {
			args[k] = &i
		}
	}
	for _, k := range []string{beforeField, afterField} {
		v, ok := args[k]
		if !ok {
			continue
		}
		c := &Cursor{}
		if c.UnmarshalGQL(v) == nil {
			args[k] = c
		}
	}
	if v, ok := args[whereField]; ok && whereInput != nil {
		if err := graphql.UnmarshalInputFromContext(ctx, v, whereInput); err == nil {
			args[whereField] = whereInput
		}
	}

	return args
}

// mayAddCondition appends another type condition to the satisfies list
// if it does not exist in the list.
func mayAddCondition(satisfies []string, typeCond []string) []string {
Cond:
	for _, c := range typeCond {
		for _, s := range satisfies {
			if c == s {
				continue Cond
			}
		}
		satisfies = append(satisfies, c)
	}
	return satisfies
}