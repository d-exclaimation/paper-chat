package graphql

import (
	"github.com/d-exclaimation/paper-chat/streaming/pubsub"
	"go.mongodb.org/mongo-driver/mongo"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	db *mongo.Database
	pubsub *pubsub.EventPubSub
}

func MakeResolver(db *mongo.Database) *Resolver {
	return &Resolver{db, pubsub.New()}
}
