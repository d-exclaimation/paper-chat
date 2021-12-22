//
//  message.go
//  model
//
//  Created by d-exclaimation on 11:12 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Message data type describing a message sent to a certain room
type Message struct {
	// Message ID
	OID primitive.ObjectID `bson:"_id"`

	// Message content / string value
	Value     string `bson:"value"`

	// Created at timestamp in ISO string format
	CreatedAt string `bson:"created_at"`

	UserID string `bson:"user_id"`

	RoomID string `bson:"room_id"`
}
