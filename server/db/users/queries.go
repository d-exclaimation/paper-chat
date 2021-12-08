//
//  queries.go
//  users
//
//  Created by d-exclaimation on 11:51 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package users

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetById(db *mongo.Database, oid primitive.ObjectID, ctx context.Context) <-chan *model.User {
	task := make(chan *model.User)

	go func() {
		var user model.User
		err := db.Collection("users").
			FindOne(ctx, bson.M{
				"_id": oid,
			}).
			Decode(&user)

		if err != nil {
			task <- nil
		} else {
			task <- &user
		}
		close(task)
	}()

	return task
}

func GetByName(db *mongo.Database, username string, ctx context.Context) <-chan *model.User {
	task := make(chan *model.User)

	go func() {
		var user model.User
		err := db.Collection("users").
			FindOne(ctx, bson.M{
				"username": username,
			}).
			Decode(&user)

		if err != nil {
			task <- nil
		} else {
			task <- &user
		}
		close(task)
	}()

	return task
}