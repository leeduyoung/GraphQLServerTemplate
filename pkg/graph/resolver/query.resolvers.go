package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.24

import (
	"context"

	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/dto"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/logger"
)

// Todos is the resolver for the todos field.
func (r *queryResolver) Todos(ctx context.Context) ([]*graphqlmodel.Todo, error) {
	response, err := r.srv.Todo.Todos(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var todoList []*graphqlmodel.Todo
	for _, v := range response {
		todoList = append(todoList, dto.NewTodoDto(v))
	}

	return todoList, nil
}

// Users is the resolver for the users field.
func (r *queryResolver) Users(ctx context.Context) ([]*graphqlmodel.User, error) {
	response, err := r.srv.User.Users(ctx)
	if err != nil {
		logger.Error(err)
		return nil, err
	}

	var userList []*graphqlmodel.User
	for _, v := range response {
		userList = append(userList, dto.NewUserDto(v))
	}

	return userList, nil
}

// Query returns gen.QueryResolver implementation.
func (r *Resolver) Query() gen.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
