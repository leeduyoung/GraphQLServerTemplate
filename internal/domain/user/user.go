package user

import "github.com/leeduyoung/GraphQLServerTemplate/ent"

type Model struct {
	ent.User
}

func New(user ent.User) *Model {
	return &Model{
		User: user,
	}
}
