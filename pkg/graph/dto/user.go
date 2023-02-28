package dto

import (
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/user"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
)

type UserDto struct {
	Name string
}

func NewUserDto(user *user.Model) *graphqlmodel.User {
	if user == nil {
		return nil
	}

	return &graphqlmodel.User{
		ID:   user.ID,
		Name: user.Name,
	}
}
