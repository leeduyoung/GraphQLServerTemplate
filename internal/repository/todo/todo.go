package todo

import (
	"context"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/ent/todo"
)

type CreateArgs struct {
	Text   string
	Done   bool
	UserID string
}

func (r repository) Create(ctx context.Context, conn *ent.TodoClient, args CreateArgs) (*ent.Todo, error) {
	return conn.Create().
		SetText(args.Text).
		SetDone(args.Done).
		SetUserID(args.UserID).
		Save(ctx)
}

type UpdateArgs struct {
	ID   int
	Text *string
	Done *bool
}

func (r repository) Update(ctx context.Context, conn *ent.TodoClient, args UpdateArgs) (*ent.Todo, error) {
	return conn.UpdateOneID(args.ID).
		SetNillableText(args.Text).
		SetNillableDone(args.Done).
		Save(ctx)
}

func (r repository) Delete(ctx context.Context, conn *ent.TodoClient, id int) error {
	return conn.DeleteOneID(id).Exec(ctx)
}

func (r repository) FindOne(ctx context.Context, conn *ent.TodoClient, id int) (*ent.Todo, error) {
	return conn.Query().
		Where(todo.ID(id)).
		Only(ctx)
}

type FindAllArgs struct {
}

func (r repository) FindAll(ctx context.Context, conn *ent.TodoClient, args FindAllArgs) ([]*ent.Todo, int, error) {
	todos, err := conn.Query().All(ctx)
	if err != nil {
		return nil, 0, err
	}

	return todos, len(todos), nil
}
