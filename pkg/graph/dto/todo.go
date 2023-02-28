package dto

import (
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/utils"
)

type TodoDto struct {
	Name string
}

func NewTodoDto(todo *todo.Model) *graphqlmodel.Todo {
	if todo == nil {
		return nil
	}

	return &graphqlmodel.Todo{
		ID:     utils.ParseIntToString(todo.ID),
		Text:   todo.Text,
		Done:   todo.Done,
		UserID: todo.UserID,
		User:   nil,
	}
}
