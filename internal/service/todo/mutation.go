package todo

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/todo"
	todoRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
)

func (s service) CreateTodo(ctx context.Context, input graphqlmodel.NewTodo) (*todo.Model, error) {
	response, err := s.todoRepository.Create(ctx, s.conn.Todo, todoRepo.CreateArgs{
		Text:   input.Text,
		Done:   false,
		UserID: input.UserID,
	})
	if err != nil {
		return nil, err
	}

	return todo.New(*response), nil
}
