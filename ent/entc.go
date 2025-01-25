//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(entgql.WithWhereInputs(true))
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
