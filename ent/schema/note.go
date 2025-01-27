package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

// Note holds the schema definition for the Note entity.
type Note struct {
	ent.Schema
}

// Fields of the Note.
func (Note) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			MinLen(3).
			Annotations(
				entgql.OrderField("TITLE"),
			),
		field.Text("body"),

		createdAtField(),
		updatedAtField(),
	}
}

// Edges of the Note.
func (Note) Edges() []ent.Edge {
	return nil
}

func (Note) Annotations() []schema.Annotation {
	return append(baseGqlAnnotations)
}
