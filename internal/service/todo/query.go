package todo

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/todo"
	todoRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/todo"
)

func (s service) Todos(ctx context.Context) ([]*todo.Model, error) {
	response, _, err := s.todoRepository.FindAll(ctx, s.conn.Todo, todoRepo.FindAllArgs{})
	if err != nil {
		return nil, err
	}

	var todoList []*todo.Model
	for _, v := range response {
		todoList = append(todoList, todo.New(*v))
	}

	return todoList, nil
}
