package todo

import "github.com/leeduyoung/GraphQLServerTemplate/ent"

type Model struct {
	ent.Todo
}

func New(todo ent.Todo) *Model {
	return &Model{
		Todo: todo,
	}
}
