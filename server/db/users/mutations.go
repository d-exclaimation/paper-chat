//
//  mutations.go
//  users
//
//  Created by d-exclaimation on 2:31 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package users

import (
	"context"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(db *mongo.Database, username string, ctx context.Context) <-chan model.SignUp {
	task := make(chan model.SignUp)
	go func() {
		var (
			user model.User
			col  = db.Collection("users")
		)

		res, err := col.InsertOne(ctx, bson.M{
			"username": username,
		})

		if err != nil {
			task <- &model.InvalidUser{Username: username, Reason: err.Error()}
			close(task)
			return
		}

		err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&user)
		if err != nil {
			task <- &model.InvalidUser{Username: username, Reason: err.Error()}
		} else {
			task <- &model.Credentials{User: &user}
		}
		close(task)
	}()
	return task
}
