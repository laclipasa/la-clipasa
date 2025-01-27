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
	deletedAtDirectiveField := DirectiveField{
		FieldName: "deletedAt",
		Targets:   []DirectiveTarget{CreateInputFieldTarget, UpdateInputFieldTarget},
		Directives: []entgql.Directive{
			schema.HasRoleDirective(user.RoleADMIN),
		},
	}
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
					Targets: []DirectiveTarget{CreateInputObjectTarget, UpdateInputObjectTarget},
					Directives: []entgql.Directive{
						schema.HasRoleDirective(user.RoleMODERATOR),
					},
				},
				{
					FieldName: "role",
					Targets:   []DirectiveTarget{CreateInputFieldTarget, UpdateInputFieldTarget},
					Directives: []entgql.Directive{
						schema.HasRoleDirective(user.RoleADMIN),
					},
				},
				deletedAtDirectiveField,
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
		log.Fatalf("failed ent codegen: %v", err)
	}
}

type DirectiveTarget int

const (
	// Field targets
	TypeFieldTarget DirectiveTarget = iota
	CreateInputFieldTarget
	UpdateInputFieldTarget

	// Object targets
	TypeObjectTarget
	CreateInputObjectTarget
	UpdateInputObjectTarget
)

func HasMixedTargets(targets []DirectiveTarget) bool {
	var combinedBits uint

	for _, target := range targets {
		combinedBits |= 1 << target
	}

	hasFieldTargets := (combinedBits & FieldTargetsMask) != 0
	hasObjectTargets := (combinedBits & ObjectTargetsMask) != 0

	return hasFieldTargets && hasObjectTargets
}

const (
	FieldTargetsMask  = 1<<TypeFieldTarget | 1<<CreateInputFieldTarget | 1<<UpdateInputFieldTarget
	ObjectTargetsMask = 1<<TypeObjectTarget | 1<<CreateInputObjectTarget | 1<<UpdateInputObjectTarget
)

type DirectiveField struct {
	// database field name
	FieldName  string
	Targets    []DirectiveTarget
	Directives []entgql.Directive
}

func applyDirectives(entityDirectives map[string][]DirectiveField) entgql.SchemaHook {
	return func(graph *gen.Graph, s *ast.Schema) error {
		for typeName, directives := range entityDirectives {
			for dfIdx, df := range directives {
				if HasMixedTargets(df.Targets) {
					return fmt.Errorf("directive (%d) for %q has mixed targets", dfIdx, typeName)
				}
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
					var genType *gen.Type
					for _, element := range graph.Nodes {
						if element.Name == typeName {
							genType = element
							break
						}
					}

					if df.FieldName != "" {
						// assign to field
						if err := addDirectiveToField(t, df.FieldName, df.Directives); err != nil {
							return fmt.Errorf("couldn't add directive to %q: %w", gqlType, err)
						}

						// extra directives may be required. see vendor/entgo.io/contrib/entgql/schema.go
						desc := entgql.MutationDescriptor{
							Type:     genType,
							IsCreate: target == CreateInputFieldTarget || target == CreateInputObjectTarget,
						}
						inputFields, err := desc.InputFields()
						if err != nil {
							return fmt.Errorf("entgql.MutationDescriptor.InputFields: %w", err)
						}
						for _, ifield := range inputFields {
							pascalDfName := gen.Field{Name: df.FieldName}.StructField()
							if ifield.StructField() != pascalDfName {
								continue
							}

							if ifield.AppendOp {
								if err := addDirectiveToField(t, "append"+ifield.StructField(), df.Directives); err != nil {
									return fmt.Errorf("couldn't add append directive to %s.%s: %w", gqlType, df.FieldName, err)
								}
							}
							if ifield.ClearOp {
								if err := addDirectiveToField(t, "clear"+ifield.StructField(), df.Directives); err != nil {
									return fmt.Errorf("couldn't add clear directive to %s.%s: %w", gqlType, df.FieldName, err)
								}
							}
						}
					} else {
						// assign to object
						t.Directives = append(t.Directives, convertDirectives(df.Directives)...)
					}

				}
			}
		}
		return nil
	}
}

func convertDirectives(directives []entgql.Directive) []*ast.Directive {
	result := make([]*ast.Directive, len(directives))
	for i, d := range directives {
		result[i] = &ast.Directive{
			Name:      d.Name,
			Arguments: d.Arguments,
		}
	}
	return result
}

func addDirectiveToField(t *ast.Definition, fieldName string, directives []entgql.Directive) error {
	field := t.Fields.ForName(fieldName)
	if field == nil {
		return fmt.Errorf("field %q not found", fieldName)
	}
	field.Directives = append(field.Directives, convertDirectives(directives)...)
	return nil
}
