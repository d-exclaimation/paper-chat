//
//  queries.go
//  messages
//
//  Created by d-exclaimation on 4:48 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package messages

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetByRoom(db *mongo.Database, roomOID primitive.ObjectID, ctx context.Context) <-chan *model.Message {
	task := make(chan *model.Message)
	go func() {
		var messages []model.Message
		find, err := db.Collection("messages").
			Find(ctx, bson.M{
				"room_id": roomOID.Hex(),
			})
		if err != nil {
			close(task)
			return
		}
		err = find.All(ctx, messages)
		if err != nil {
			close(task)
			return
		}

		for _, message := range messages {
			task <- &message
		}
		close(task)
	}()
	return task
}