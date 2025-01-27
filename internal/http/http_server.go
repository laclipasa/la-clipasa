package http

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"entgo.io/ent/dialect"
	entsql "entgo.io/ent/dialect/sql"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"

	"github.com/go-chi/chi/v5/middleware"
	"github.com/laclipasa/la-clipasa/ent"
	"github.com/laclipasa/la-clipasa/gql"
	"github.com/laclipasa/la-clipasa/internal"
	"github.com/laclipasa/la-clipasa/internal/envvar"
	postgresql "github.com/laclipasa/la-clipasa/internal/postgres"
)

/***
 *
 * TODO: SSE via gqlgen: https://gqlgen.com/recipes/subscriptions/#adding-server-sent-events-transport
 *  				 and https://github.com/enisdenjo/graphql-sse
 */

// Open new connection
func Open(databaseUrl string) *ent.Client {
	db, err := sql.Open("pgx", databaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	// Create an ent.Driver from `db`.
	drv := entsql.OpenDB(dialect.Postgres, db)
	return ent.NewClient(ent.Driver(drv))
}

func RunServer(envFile string) {
	if err := envvar.Load(envFile); err != nil {
		fmt.Printf("envvar.Load: %s\n", err)
		return
	}
	if err := internal.NewAppConfig(); err != nil {
		fmt.Printf("NewAppConfig: %s\n", err)
		return
	}

	cfg := internal.Config
	logger, err := zap.NewProduction()
	if err != nil {
		fmt.Printf("zap.NewProduction: %s\n", err)
		return
	}
	defer logger.Sync()

	db, sqldb, err := postgresql.New(logger.Sugar(), postgresql.WithDBName(cfg.Postgres.DB))
	if err != nil {
		fmt.Printf("postgresql.New: %s\n", err)
		return
	}
	defer db.Close()

	drv := entsql.OpenDB(dialect.Postgres, sqldb)

	entClient := ent.NewClient(ent.Driver(drv))
	defer entClient.Close()

	ctx := ent.NewContext(context.Background(), entClient)

	/*
		this automagically syns the schema with the database. do not use
		if err := entClient.Schema.Create(ctx); err != nil {
			log.Fatalf("failed initializing migrations table: %v", err)
		} */

	router := chi.NewRouter()

	router.Use(
		middleware.RequestID,
		middleware.RealIP,
		middleware.Logger,
		middleware.Recoverer,
	)

	srv := gql.NewServer(ctx)
	router.Handle("/graphql", srv)
	router.Handle("/graphiql", playground.ApolloSandboxHandler("GraphQL", "/graphql"))

	addr := "localhost:" + cfg.APIPort
	fmt.Printf("starting server at :%s\n", addr)
	fmt.Printf("playground: http://localhost:%s/graphiql\n", cfg.APIPort)
	if err := ListenAndServe(ctx, addr, router, 5*time.Second); err != nil {
		log.Fatalf("ListenAndServe: %s\n", err)
	}
}

// TODO: migrate internal/rest.NewServer instead
// for internal/rest.runTestServer hopefully an adaptation of https://github.com/99designs/gqlgen/blob/master/client/client.go
// or using https://github.com/Khan/genqlient suffices
func ListenAndServe(ctx context.Context, addr string, handler http.Handler, shutdownTimeout time.Duration) error {
	server := &http.Server{Addr: addr, Handler: handler}

	errChan := make(chan error)
	go func() {
		if err := server.ListenAndServe(); err != nil {
			errChan <- err
		}
	}()

	select {
	case err := <-errChan:
		return err

	case <-ctx.Done():
		// do nothing
	}

	ctx, cancel := context.WithTimeout(context.Background(), shutdownTimeout)
	defer cancel()

	return server.Shutdown(ctx)
}
