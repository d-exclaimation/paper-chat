//
//  connect.go
//  db
//
//  Created by d-exclaimation on 8:49 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package db

import (
	"context"
	"github.com/d-exclaimation/paper-chat/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Mongo struct {
	ctx context.Context
	client *mongo.Client
}

func MakeMongo() (*Mongo, error){
	ctx := context.Background()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(config.MongoURI()))
	if err != nil {
		return nil, err
	}
	return &Mongo{
		ctx:    ctx,
		client: client,
	}, nil
}

func (m *Mongo) Database(name string, opt ...*options.DatabaseOptions) *mongo.Database {
	return m.client.Database(name, opt...)
}

func (m *Mongo) Disconnect() {
	if err := m.client.Disconnect(m.ctx); err != nil {
		log.Println(err)
	}
}
