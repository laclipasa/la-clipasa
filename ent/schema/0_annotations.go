package schema

import (
	"entgo.io/contrib/entgql"
	"entgo.io/ent/schema"
)

var baseGqlAnnotations = []schema.Annotation{
	entgql.QueryField(),
	entgql.RelayConnection(),
	entgql.Directives(),
	entgql.Mutations(entgql.MutationCreate(), entgql.MutationUpdate()),
}
