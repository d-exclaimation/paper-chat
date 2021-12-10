package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/d-exclaimation/paper-chat/graphql/gql"
)

func (r *mutationResolver) Hello(_ context.Context, msg string) (string, error) {
	return msg, nil
}

func (r *queryResolver) Hello(_ context.Context) (string, error) {
	return "", nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
