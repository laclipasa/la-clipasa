//go:build ignore
// +build ignore

package main

import (
	"fmt"
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/laclipasa/la-clipasa/ent/schema"
	"github.com/laclipasa/la-clipasa/ent/user"
	"github.com/vektah/gqlparser/v2/ast"
)

func main() {
	ex, err := entgql.NewExtension(
		// Tell Ent to generate a GraphQL schema for
		// the Ent schema in a file named ent.graphql.
		entgql.WithSchemaGenerator(),
		entgql.WithWhereInputs(true),
		// required for extra gen
		entgql.WithConfigPath("../gqlgen.yml"),
		entgql.WithSchemaPath("../gql/schema/ent.graphql"),
		entgql.WithSchemaHook(applyDirectives(map[string][]DirectiveField{
			"User": {
				{
					FieldName: "role",
					Targets:   []DirectiveTarget{CreateInputFieldTarget, UpdateInputFieldTarget},
					Directives: []entgql.Directive{
						schema.HasRoleDirective(user.RoleADMIN),
					},
				},
			},
		})),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}
	log.Default().Printf("Running ent codegen...")
	if err := entc.Generate("./schema", &gen.Config{
		Features: []gen.Feature{
			gen.FeatureVersionedMigration,
			gen.FeatureLock,
			gen.FeatureIntercept,
			gen.FeatureSnapshot,
		},
	}, entc.Extensions(ex)); err != nil {
		log.Fatalf("running ent codegen: %v", err)
	}
}

type DirectiveTarget int

const (
	TypeFieldTarget DirectiveTarget = iota
	CreateInputFieldTarget
	UpdateInputFieldTarget

	TypeObjectTarget
	CreateInputObjectTarget
	UpdateInputObjectTarget
)

type DirectiveField struct {
	FieldName  string
	Targets    []DirectiveTarget
	Directives []entgql.Directive
}

func applyDirectives(entityDirectives map[string][]DirectiveField) entgql.SchemaHook {
	return func(_ *gen.Graph, s *ast.Schema) error {
		for typeName, directives := range entityDirectives {
			for _, df := range directives {
				for _, target := range df.Targets {
					var gqlType string
					switch target {
					case TypeFieldTarget, TypeObjectTarget:
						gqlType = typeName
					case CreateInputFieldTarget, CreateInputObjectTarget:
						gqlType = "Create" + typeName + "Input"
					case UpdateInputFieldTarget, UpdateInputObjectTarget:
						gqlType = "Update" + typeName + "Input"
					}

					t := s.Types[gqlType]
					if t == nil {
						return fmt.Errorf("type %q not found", gqlType)
					}

					if df.FieldName != "" {
						// assume directive assigned to field
						f := t.Fields.ForName(df.FieldName)
						if f == nil {
							return fmt.Errorf("field %q not found in type %q", df.FieldName, gqlType)
						}
						for _, d := range df.Directives {
							f.Directives = append(f.Directives, &ast.Directive{
								Name:      d.Name,
								Arguments: d.Arguments,
							})
						}
					} else {
						// assume directive assigned to object
						for _, d := range df.Directives {
							t.Directives = append(t.Directives, &ast.Directive{
								Name:      d.Name,
								Arguments: d.Arguments,
							})
						}
					}

				}
			}
		}
		return nil
	}
}
