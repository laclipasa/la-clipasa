package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("display_name"),
		field.String("profile_image").
			Nillable().
			Optional(),
		field.String("twitch_id").
			Unique(),
		field.Enum("role").
			Values("USER", "ADMIN", "MODERATOR").
			Default("USER"),
		field.JSON("awards", []string{}).
			Optional(),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Time("deleted_at").
			Optional().
			Nillable(),
	}
}

func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("saved_posts", Post.Type),
		edge.To("liked_posts", Post.Type),
		edge.To("posts", Post.Type),
		edge.To("comments", Comment.Type),
	}
}

func (User) Annotations() []schema.Annotation {
	return append(baseGqlAnnotations)
}
