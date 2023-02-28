package service

import (
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/service/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/service/user"
)

// ServiceRoot ...
type ServiceRoot struct {
	User user.Service
	Todo todo.Service
}

// New ...
func New(conn *ent.Client, repository *repository.RepositoryRoot) ServiceRoot {
	return ServiceRoot{
		User: user.New(conn, repository),
		Todo: todo.New(conn, repository),
	}
}
