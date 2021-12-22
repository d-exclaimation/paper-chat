package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"time"

	"github.com/d-exclaimation/paper-chat/graphql/gql"
)

func (r *mutationResolver) Hello(ctx context.Context, msg string) (string, error) {
	return msg, nil
}

func (r *queryResolver) Hello(ctx context.Context) (string, error) {
	return "", nil
}

func (r *subscriptionResolver) Hello(ctx context.Context) (<-chan string, error) {
	res := make(chan string)
	go func() {
		res <- "Hello"
		time.Sleep(1 * time.Second)
		close(res)
	}()
	return res, nil
}

// Mutation returns gql.MutationResolver implementation.
func (r *Resolver) Mutation() gql.MutationResolver { return &mutationResolver{r} }

// Query returns gql.QueryResolver implementation.
func (r *Resolver) Query() gql.QueryResolver { return &queryResolver{r} }

// Subscription returns gql.SubscriptionResolver implementation.
func (r *Resolver) Subscription() gql.SubscriptionResolver { return &subscriptionResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type subscriptionResolver struct{ *Resolver }
