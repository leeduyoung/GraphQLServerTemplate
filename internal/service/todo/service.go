package todo

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository"
	todoRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
)

type Service interface {
	Todos(ctx context.Context) ([]*todo.Model, error)
	CreateTodo(ctx context.Context, input graphqlmodel.NewTodo) (*todo.Model, error)
}

type service struct {
	conn           *ent.Client
	todoRepository todoRepo.Repository
}

func New(conn *ent.Client, repository *repository.RepositoryRoot) Service {
	return &service{
		conn:           conn,
		todoRepository: repository.Todo,
	}
}
