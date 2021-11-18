package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"go.mongodb.org/mongo-driver/bson/primitive"

	"github.com/d-exclaimation/paper-chat/graphql/gql"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
)

func (r *queryResolver) RandomRoom(ctx context.Context) (*model.Room, error) {
	var (
		room model.Room
		doc  = r.db.Collection("rooms").FindOne(ctx, bson.M{})
	)
	if err := doc.Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *queryResolver) Room(ctx context.Context, id string) (*model.Room, error) {
	if !primitive.IsValidObjectID(id) {
		return nil, nil
	}

	var (
		room model.Room
		oid, _ = primitive.ObjectIDFromHex(id)
		query  = bson.M{
			"_id": oid,
		}
		doc = r.db.Collection("rooms").FindOne(ctx, query)
	)
	if err := doc.Decode(&room); err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *roomResolver) ID(ctx context.Context, obj *model.Room) (string, error) {
	return obj.OID.Hex(), nil
}

// Room returns gql.RoomResolver implementation.
func (r *Resolver) Room() gql.RoomResolver { return &roomResolver{r} }

type roomResolver struct{ *Resolver }
