package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/d-exclaimation/paper-chat/db/rooms"
	"github.com/d-exclaimation/paper-chat/graphql/gql"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (r *mutationResolver) CreateRoom(ctx context.Context, title string) (model.CreateResult, error) {
	return <-rooms.New(r.db, title, ctx), nil
}

func (r *mutationResolver) JoinRoom(ctx context.Context, id string) (model.JoinResult, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return &model.RoomDoesntExist{ID: id}, nil
	}
	user := &model.User{
		OID: primitive.NewObjectID(),
		Username: "Vincent",
	}
	return <-rooms.Join(r.db, oid, user, ctx), nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, nil
	}

	return <-rooms.GetById(r.db, oid, ctx), nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return obj.OID.Hex(), nil
}

// Room returns gql.RoomResolver implementation.
func (r *Resolver) Room() gql.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
