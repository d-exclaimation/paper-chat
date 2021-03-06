package graphql

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"errors"
	"github.com/d-exclaimation/paper-chat/db/messages"
	"github.com/d-exclaimation/paper-chat/db/rooms"
	"github.com/d-exclaimation/paper-chat/graphql/auth"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"

	"github.com/d-exclaimation/paper-chat/graphql/model"
)

func (r *mutationResolver) Send(ctx context.Context, roomID string, content string) (model.SendResult, error) {
	oid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return model.RoomDoesntExist{ID: roomID}, nil
	}

	room := <-rooms.GetById(r.db, oid, ctx)
	if room == nil {
		return model.RoomDoesntExist{ID: roomID}, nil
	}

	user := <-auth.Auth(r.db, ctx)
	if user == nil {
		return model.NotLoggedIn{Username: nil}, nil
	}

	isParticipant := false
	for _, participant := range room.Participant {
		if participant.OID.Hex() == user.OID.Hex() && participant.Username == user.Username {
			isParticipant = true
			break
		}
	}

	if !isParticipant {
		return model.NotAParticipant{ID: user.OID.Hex(), Username: user.Username}, nil
	}

	// TODO: Save to database
	msg := &model.Message{
		OID:       primitive.NewObjectID(),
		Value:     content,
		CreatedAt: time.Now().UTC().String(),
		UserID:    user.OID.Hex(),
		RoomID:    roomID,
	}

	r.pubsub.Publish(roomID, msg)

	return model.SendSuccessful{Message: msg}, nil
}

func (r *subscriptionResolver) Chat(ctx context.Context, roomID string) (<-chan *model.Message, error) {
	oid, err := primitive.ObjectIDFromHex(roomID)
	if err != nil {
		return nil, err
	}

	room := <-rooms.GetById(r.db, oid, ctx)
	if room == nil {
		return nil, errors.New("room doesn't exist")
	}

	messagesFuture := messages.GetByRoom(r.db, oid, ctx)

	channel := r.pubsub.Stream(roomID, ctx)

	go func() {
		msgs := <-messagesFuture
		if msgs == nil {
			return
		}
		for _, message := range msgs {
			channel <- &message
		}
	}()

	return channel, nil
}
