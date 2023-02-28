package storage

// import graph gophers with your other imports
import (
	"context"
	"fmt"
	"github.com/graph-gophers/dataloader"
	"github.com/leeduyoung/GraphQLServerTemplate/ent"
	"github.com/leeduyoung/GraphQLServerTemplate/ent/user"
	userModel "github.com/leeduyoung/GraphQLServerTemplate/internal/domain/user"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/dto"
	"github.com/leeduyoung/GraphQLServerTemplate/pkg/graph/gen/graphqlmodel"
	"net/http"
)

type ctxKey string

const (
	loadersKey = ctxKey("dataloaders")
)

// UserReader reads Users from a database
type UserReader struct {
	conn *ent.Client
}

// GetUsers implements a batch function that can retrieve many users by ID,
// for use in a dataloader
func (u *UserReader) GetUsers(ctx context.Context, keys dataloader.Keys) []*dataloader.Result {
	// read all requested users in a single query
	userIDs := make([]string, len(keys))
	for ix, key := range keys {
		userIDs[ix] = key.String()
	}

	res, err := u.conn.User.Query().Where(user.IDIn(userIDs...)).All(ctx)
	if err != nil {
		panic(err)
	}

	userByID := map[string]*graphqlmodel.User{}
	for _, v := range res {
		model := userModel.New(*v)
		userByID[v.ID] = dto.NewUserDto(model)
	}

	output := make([]*dataloader.Result, len(keys))
	for i, userKey := range keys {
		user, ok := userByID[userKey.String()]
		if ok {
			output[i] = &dataloader.Result{
				Data:  user,
				Error: nil,
			}
		} else {
			err := fmt.Errorf("user not found %s", userKey.String())
			output[i] = &dataloader.Result{
				Data:  nil,
				Error: err,
			}
		}
	}

	return output
}

// Loaders wrap your data loaders to inject via middleware
type Loaders struct {
	UserLoader *dataloader.Loader
}

// NewLoaders instantiates data loaders for the middleware
//func NewLoaders(conn *sql.DB) *Loaders {
func NewLoaders(conn *ent.Client) *Loaders {
	// define the data loader
	userReader := &UserReader{conn: conn}
	loaders := &Loaders{
		UserLoader: dataloader.NewBatchedLoader(userReader.GetUsers),
	}
	return loaders
}

// Middleware injects data loaders into the context
func Middleware(loaders *Loaders, next http.Handler) http.Handler {
	// return a middleware that injects the loader to the request context
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		nextCtx := context.WithValue(r.Context(), loadersKey, loaders)
		r = r.WithContext(nextCtx)
		next.ServeHTTP(w, r)
	})
}

// For returns the dataloader for a given context
func For(ctx context.Context) *Loaders {
	return ctx.Value(loadersKey).(*Loaders)
}

// GetUser wraps the User dataloader for efficient retrieval by user ID
func GetUser(ctx context.Context, userID string) (*graphqlmodel.User, error) {
	loaders := For(ctx)
	thunk := loaders.UserLoader.Load(ctx, dataloader.StringKey(userID))
	result, err := thunk()
	if err != nil {
		return nil, err
	}
	return result.(*graphqlmodel.User), nil
}
