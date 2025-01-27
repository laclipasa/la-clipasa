package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Comment holds the schema definition for the Comment entity.
type Comment struct {
	ent.Schema
}

// Fields of the Comment.
func (Comment) Fields() []ent.Field {
	return []ent.Field{
		field.String("content").
			NotEmpty(),

		createdAtField(),
		updatedAtField(),
		deletedAtField(),
	}
}

// Edges of the Comment.
func (Comment) Edges() []ent.Edge {
	return []ent.Edge{
		// Comment belongs to an author only
		edge.From("author", User.Type).
			Ref("comments").
			Unique(),
		// Comment belongs to a post only
		edge.From("post", Post.Type).
			Ref("comments").
			Unique(),
	}
}

func (Comment) Annotations() []schema.Annotation {
	return append(baseGqlAnnotations)
}
