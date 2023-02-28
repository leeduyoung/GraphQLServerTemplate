package todo

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
)

// Repository ...
type Repository interface {
	Create(ctx context.Context, conn *ent.TodoClient, args CreateArgs) (*ent.Todo, error)
	Update(ctx context.Context, conn *ent.TodoClient, args UpdateArgs) (*ent.Todo, error)
	Delete(ctx context.Context, conn *ent.TodoClient, todoID int) error
	FindOne(ctx context.Context, conn *ent.TodoClient, todoID int) (*ent.Todo, error)
	FindAll(ctx context.Context, conn *ent.TodoClient, args FindAllArgs) ([]*ent.Todo, int, error)
}

type repository struct {
}

// New ...
func New() Repository {
	return &repository{}
}
