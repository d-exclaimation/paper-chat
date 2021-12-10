package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/d-exclaimation/paper-chat/db/users"
	"github.com/d-exclaimation/paper-chat/graphql/auth"
	"github.com/d-exclaimation/paper-chat/graphql/gql"
	"github.com/d-exclaimation/paper-chat/graphql/model"
)

func (r *mutationResolver) Signup(ctx context.Context, username string) (model.SignUp, error) {
	res := <-users.Create(r.db, username, ctx)
	switch res.(type) {
	case *model.Credentials:
		user := res.(*model.Credentials).User
		token, exp, err := auth.Sign(user)
		if err != nil {
			return &model.InvalidUser{Username: username, Reason: err.Error()}, nil
		}
		return &model.Credentials{
			AccessToken: token,
			ExpireAt:    fmt.Sprintf("%s", exp.UTC().String()),
			User:        user,
		}, nil
	default:
		return res, nil
	}
}

func (r *mutationResolver) Login(ctx context.Context, username string) (model.LogIn, error) {
	user := <-users.GetByName(r.db, username, ctx)
	if user == nil {
		return &model.InvalidUser{Username: username, Reason: "Cannot find user"}, nil
	}

	token, exp, err := auth.Sign(user)
	if err != nil {
		return &model.InvalidUser{Username: username, Reason: err.Error()}, nil
	}

	return &model.Credentials{
		AccessToken: token,
		ExpireAt:    fmt.Sprintf("%s", exp.UTC().String()),
		User:        user,
	}, nil
}

func (r *queryResolver) Me(ctx context.Context) (*model.User, error) {
	return <-auth.Auth(r.db, ctx), nil
}

func (r *userResolver) ID(_ context.Context, obj *model.User) (string, error) {
	return obj.OID.Hex(), nil
}

// User returns gql.UserResolver implementation.
func (r *Resolver) User() gql.UserResolver { return &userResolver{r} }

type userResolver struct{ *Resolver }
