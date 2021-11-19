//
//  mutations.go
//  db
//
//  Created by d-exclaimation on 8:43 AM.
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

type RoomResult struct {
	Room  *model.Room
	Error error
}

func New(db *mongo.Database, title string, ctx context.Context) <-chan RoomResult {
	task := make(chan RoomResult)
	go func() {
		var (
			room model.Room
			col  = db.Collection("rooms")
		)
		res, err := col.InsertOne(ctx, bson.M{
			"title":       title,
			"participant": []bson.M{},
		})
		if err != nil {
			task <- RoomResult{Error: err}
			return
		}
		err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&room)
		if err != nil {
			task <- RoomResult{Error: err}
		} else {
			task <- RoomResult{Room: &room}
		}
		close(task)
	}()
	return task
}

func Join(db *mongo.Database, oid primitive.ObjectID, user *model.User, ctx context.Context) <-chan RoomResult {
	task := make(chan RoomResult)
	go func() {
		var (
			col            = db.Collection("rooms")
			initial        = <-GetById(db, oid, ctx)
			newParticipant = append(initial.Participant, user)
			update         = bson.M{"$set": bson.M{"participant": newParticipant}}
		)
		_, err := col.UpdateOne(ctx, bson.M{}, update)
		if err != nil {
			task <- RoomResult{Error: err}
		} else {
			task <- RoomResult{Room: &model.Room{OID: oid, Title: initial.Title, Participant: newParticipant}}
		}
	}()
	return task
}
