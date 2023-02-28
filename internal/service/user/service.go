package user

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/domain/user"
	"github.com/leeduyoung/GraphQLServerTemplate/internal/repository"
	userRepo "github.com/leeduyoung/GraphQLServerTemplate/internal/repository/user"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
)

type Service interface {
	Users(ctx context.Context) ([]*user.Model, error)
	CreateUser(ctx context.Context, input graphqlmodel.NewUser) (*user.Model, error)
}

type service struct {
	conn           *ent.Client
	userRepository userRepo.Repository
}

func New(conn *ent.Client, repository *repository.RepositoryRoot) Service {
	return &service{
		conn:           conn,
		userRepository: repository.User,
	}
}
