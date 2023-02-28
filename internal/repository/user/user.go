package user

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/ent/user"
)

type CreateArgs struct {
	Name string
}

func (r repository) Create(ctx context.Context, conn *ent.UserClient, args CreateArgs) (*ent.User, error) {
	return conn.Create().
		SetName(args.Name).
		Save(ctx)
}

type UpdateArgs struct {
	ID   string
	Name *string
}

func (r repository) Update(ctx context.Context, conn *ent.UserClient, args UpdateArgs) (*ent.User, error) {
	return conn.UpdateOneID(args.ID).
		SetNillableName(args.Name).
		Save(ctx)
}

func (r repository) Delete(ctx context.Context, conn *ent.UserClient, id string) error {
	return conn.DeleteOneID(id).
		Exec(ctx)
}

func (r repository) FindOne(ctx context.Context, conn *ent.UserClient, id string) (*ent.User, error) {
	return conn.Query().
		Where(user.ID(id)).
		Only(ctx)
}

type FindAllArgs struct {
}

func (r repository) FindAll(ctx context.Context, conn *ent.UserClient, args FindAllArgs) ([]*ent.User, int, error) {
	users, err := conn.Query().All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return users, len(users), nil
}
