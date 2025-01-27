package schema

import (
	"time"

	"entgo.io/contrib/entgql"
	"entgo.io/ent"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
	"github.com/laclipasa/la-clipasa/ent/user"
	"github.com/vektah/gqlparser/v2/ast"
)

var baseGqlAnnotations = []schema.Annotation{
	entgql.QueryField(),
	entgql.RelayConnection(),
	entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
}

func HasRoleDirective(role user.Role) entgql.Directive {
	var args []*ast.Argument
	if role != "" {
		args = append(args, &ast.Argument{
			Name: "role",
			Value: &ast.Value{
				Raw:  string(role),
				Kind: ast.EnumValue,
			},
		})
	}
	return entgql.NewDirective("hasRole", args...)
}

func updatedAtField() ent.Field {
	return field.Time("updated_at").
		Default(time.Now).
		UpdateDefault(time.Now).
		Annotations(
			entgql.OrderField("UPDATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		)
}

func createdAtField() ent.Field {
	return field.Time("created_at").
		Default(time.Now).
		Immutable().
		Annotations(
			entgql.OrderField("CREATED_AT"),
			entgql.Skip(entgql.SkipMutationCreateInput, entgql.SkipMutationUpdateInput),
		)
}

func deletedAtField() ent.Field {
	return field.Time("deleted_at").
		Optional().
		Nillable()
}
