package resolver

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/service"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	ctx  context.Context
	conn *ent.Client
	srv  service.ServiceRoot
	repo *repository.RepositoryRoot
}

func New(ctx context.Context, conn *ent.Client, srv service.ServiceRoot,
	repo *repository.RepositoryRoot) *Resolver {
	return &Resolver{
		ctx:  ctx,
		conn: conn,
		srv:  srv,
		repo: repo,
	}
}
