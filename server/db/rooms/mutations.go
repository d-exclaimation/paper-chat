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

func New(db *mongo.Database, title string, ctx context.Context) <-chan model.CreateResult {
	task := make(chan model.CreateResult)
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
			task <- &model.OperationFailed{Reason: err.Error()}
			close(task)
			return
		}
		err = col.FindOne(ctx, bson.M{"_id": res.InsertedID}).Decode(&room)
		if err != nil {
			task <- &model.OperationFailed{Reason: err.Error()}
		} else {
			task <- &model.RoomSuccessOperation{Payload: &room}
		}
		close(task)
	}()
	return task
}

func Join(db *mongo.Database, oid primitive.ObjectID, user *model.User, ctx context.Context) <-chan model.JoinResult {
	task := make(chan model.JoinResult)
	go func() {
		var (
			col     = db.Collection("rooms")
			initial = <-GetById(db, oid, ctx)
		)

		if initial == nil {
			task <- &model.RoomDoesntExist{ID: oid.Hex()}
			close(task)
			return
		}

		for _, p := range initial.Participant {
			if p.OID.Hex() == user.OID.Hex() {
				task <- &model.AlreadyJoined{ID: user.OID.Hex(), Username: user.Username}
				close(task)
				return
			}
		}

		newParticipant := append(initial.Participant, user)
		update := bson.M{"$set": bson.M{"participant": newParticipant}}
		_, err := col.UpdateOne(ctx, bson.M{"_id": oid}, update)
		if err != nil {
			task <- &model.OperationFailed{Reason: err.Error()}
		} else {
			task <- &model.RoomSuccessOperation{Payload: &model.Room{OID: oid, Title: initial.Title, Participant: newParticipant}}
		}
		close(task)
	}()
	return task
}

func Leave(db *mongo.Database, oid primitive.ObjectID, user *model.User, ctx context.Context) <-chan model.LeaveResult {
	task := make(chan model.LeaveResult)
	go func() {
		var (
			col     = db.Collection("rooms")
			initial = <-GetById(db, oid, ctx)
		)

		if initial == nil {
			task <- &model.RoomDoesntExist{ID: oid.Hex()}
			close(task)
			return
		}

		isIn := false
		for _, p := range initial.Participant {
			if p.OID.Hex() == user.OID.Hex() {
				isIn = true
				break
			}
		}

		if !isIn {
			task <- &model.NotAParticipant{
				ID:       user.OID.Hex(),
				Username: user.Username,
			}
			close(task)
			return
		}

		i := 0
		newParticipant := make([]*model.User, len(initial.Participant)-1)
		for _, p := range initial.Participant {
			if p.OID.Hex() != user.OID.Hex() {
				newParticipant[i] = p
				i++
			}
		}
		update := bson.M{"$set": bson.M{"participant": newParticipant}}
		_, err := col.UpdateOne(ctx, bson.M{"_id": oid}, update)
		if err != nil {
			task <- &model.OperationFailed{Reason: err.Error()}
		} else {
			task <- &model.RoomSuccessOperation{Payload: &model.Room{OID: oid, Title: initial.Title, Participant: newParticipant}}
		}
		close(task)
	}()
	return task
}
