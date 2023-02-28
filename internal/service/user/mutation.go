package user

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/user"
	userRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/user"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
)

func (s service) CreateUser(ctx context.Context, input graphqlmodel.NewUser) (*user.Model, error) {
	response, err := s.userRepository.Create(ctx, s.conn.User, userRepo.CreateArgs{
		Name: input.Name,
	})
	if err != nil {
		return nil, err
	}

	return user.New(*response), nil
}
