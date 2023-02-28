package user

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, conn *ent.UserClient, args CreateArgs) (*ent.User, error)
	Update(ctx context.Context, conn *ent.UserClient, args UpdateArgs) (*ent.User, error)
	Delete(ctx context.Context, conn *ent.UserClient, id string) error
	FindOne(ctx context.Context, conn *ent.UserClient, id string) (*ent.User, error)
	FindAll(ctx context.Context, conn *ent.UserClient, args FindAllArgs) ([]*ent.User, int, error)
}

type repository struct {
}

// NewRepository ...
func New() Repository {
	return &repository{}
}
