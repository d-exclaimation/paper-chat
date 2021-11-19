//
//  queries.go
//  db
//
//  Created by d-exclaimation on 8:30 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package rooms

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetById(db *mongo.Database, oid primitive.ObjectID, ctx context.Context) <-chan *model.Room {
	task := make(chan *model.Room)
	go func() {
		var room model.Room
		err := db.Collection("rooms").
			FindOne(ctx, bson.M{
				"_id": oid,
			}).
			Decode(&room)

		if err != nil {
			task <- nil
		} else {
			task <- &room
		}
		close(task)
	}()
	return task
}
