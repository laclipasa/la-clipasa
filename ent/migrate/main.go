//go:build ignore

package main

import (
	"context"
	"log"
	"net"
	"net/url"
	"os"

	"ariga.io/atlas/sql/sqltool"
	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"

	"github.com/laclipasa/la-clipasa/ent/migrate"
	"github.com/laclipasa/la-clipasa/internal"
	"github.com/laclipasa/la-clipasa/internal/envvar"
	_ "github.com/lib/pq"
)

// Usage: go run -mod=mod ent/migrate/main.go <name>
func main() {
	// this is used only in dev mode to generate versioned migration files.
	if err := envvar.Load(".env.dev"); err != nil {
		log.Fatalf("failed loading .env.dev file: %v", err)
	}
	if err := internal.NewAppConfig(); err != nil {
		log.Fatalf("failed loading app config: %v", err)
	}
	cfg := internal.Config

	ctx := context.Background()
	// Create a local migration directory able to understand golang-migrate migration file format for replay.
	dir, err := sqltool.NewGolangMigrateDir("ent/migrate/migrations")
	if err != nil {
		log.Fatalf("failed creating atlas migration directory: %v", err)
	}
	// Migrate diff options.
	opts := []schema.MigrateOption{
		schema.WithDir(dir),                         // provide migration directory
		schema.WithMigrationMode(schema.ModeReplay), // provide migration mode
		schema.WithDialect(dialect.Postgres),        // Ent dialect to use
	}
	if len(os.Args) != 2 {
		log.Fatalln("migration name is required. Use: 'go run -mod=mod ent/migrate/main.go <name>'")
	}

	dsn := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.Postgres.User, cfg.Postgres.Password),
		Host:   net.JoinHostPort(cfg.Postgres.Server, cfg.Postgres.Port),
		Path:   cfg.Postgres.DB,
	}

	err = migrate.NamedDiff(ctx, dsn.String()+"?sslmode=disable", os.Args[1], opts...)
	if err != nil {
		log.Fatalf("failed generating migration file: %v", err)
	}
}
