package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Post holds the schema definition for the Post entity.
type Post struct {
	ent.Schema
}

// Fields of the Post.
func (Post) Fields() []ent.Field {
	return []ent.Field{
		field.Bool("pinned").
			Default(false),
		field.String("title"),
		field.String("content").
			Nillable().
			Optional(),
		field.String("link").
			NotEmpty(),
		field.String("moderation_comment").
			Optional(),
		field.Bool("is_moderated").
			Default(false),
		field.Time("created_at").
			Default(time.Now).
			Immutable(),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
		field.Enum("categories").
			Values("RANA", "SIN_SONIDO", "MEME_ARTESANAL", "NO_SE_YO", "ORO", "DIAMANTE", "MEH", "ALERTA_GLONETILLO", "GRR", "ENSORDECEDOR", "RAGUUUL").SchemaType(map[string]string{
			dialect.Postgres: "post_categories",
		}),

		// TODO: text searchable field with gin trigrams
		// field.Other
	}
}

// Edges of the Post.
func (Post) Edges() []ent.Edge {
	return []ent.Edge{
		// Post belongs to a user (author)
		edge.From("author", User.Type).
			Ref("posts").
			Unique(),
		// Post has many comments
		edge.To("comments", Comment.Type),
		// Saved or liked posts "belong" to a user
		edge.From("saved_by", User.Type).
			Ref("saved_posts"),
		edge.From("liked_by", User.Type).
			Ref("liked_posts"),
	}
}

func (Post) Annotations() []schema.Annotation {
	return append(baseGqlAnnotations)
}
