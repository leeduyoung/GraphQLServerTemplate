package repository

import (
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository/todo"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository/user"
	"sync"
)

var (
	repository *RepositoryRoot
	once       sync.Once
)

type RepositoryRoot struct {
	User user.Repository
	Todo todo.Repository
}

func New() *RepositoryRoot {
	if repository == nil {
		once.Do(func() {
			repository = &RepositoryRoot{
				User: user.New(),
				Todo: todo.New(),
			}
		})
	}

	return repository
}
