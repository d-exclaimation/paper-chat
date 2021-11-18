package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
)

func (r *queryResolver) RandomRoom(_ context.Context) (*model.Room, error) {
	room := &model.Room{ID: "ok"}
	return room, nil
}
