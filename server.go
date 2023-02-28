package main

import (
	"context"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/service"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/config"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/db"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/resolver"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/logger"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/storage"
	"log"
	"net/http"
)

func main() {
	conn, repo, srv := initialize()

	resolvers := resolver.New(context.Background(), conn, srv, repo)
	loader := storage.NewLoaders(conn)

	svc := handler.NewDefaultServer(gen.NewExecutableSchema(
		gen.Config{
			Resolvers: resolvers,
		},
	))
	dataloaderSrv := storage.Middleware(loader, svc)
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", dataloaderSrv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", config.Instance.App.Port)
	log.Fatal(http.ListenAndServe(":"+config.Instance.App.Port, nil))
}

func initialize() (*ent.Client, *repository.RepositoryRoot, service.ServiceRoot) {
	// CREATE CONFIG
	cfg := config.Initialize()

	// CREATE LOGGER
	logger.Initialize(cfg.App.ServiceName, logger.Config{
		Mode: pkg.Mode(cfg.App.Mode),
	})

	// CREATE DATABASE
	conn := db.New(pkg.Mode(cfg.App.Mode))

	// CREATE REPOSITORY
	repo := repository.New()

	// CREATE SERVICE
	srv := service.New(conn, repo)
	return conn, repo, srv
}
