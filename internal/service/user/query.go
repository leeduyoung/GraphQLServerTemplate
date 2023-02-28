package user

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/user"
	userRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/user"
)

func (s service) Users(ctx context.Context) ([]*user.Model, error) {
	response, _, err := s.userRepository.FindAll(ctx, s.conn.User, userRepo.FindAllArgs{})
	if err != nil {
		return nil, err
	}

	var userList []*user.Model
	for _, v := range response {
		userList = append(userList, user.New(*v))
	}

	return userList, nil
}
