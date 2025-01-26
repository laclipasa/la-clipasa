package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
	"github.com/laclipasa/la-clipasa/ent/user"
	"github.com/vektah/gqlparser/v2/ast"
)

var baseGqlAnnotations = []schema.Annotation{
	entgql.QueryField(),
	entgql.RelayConnection(),
	entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
}

func hasRoleDirective(role user.Role) entgql.Directive {
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

func newHasRoleFieldInputDirective(role user.Role) entgql.Directive {
	d := hasRoleDirective(role)
	d.Skip = []entgql.SkipDirective{entgql.SkipDirectiveField}

	return d
}
