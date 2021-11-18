//
//  room.go
//  model
//
//  Created by d-exclaimation on 9:48 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// Room data type describing metadata for a certain chat group
type Room struct {
	// Room ID
	OID primitive.ObjectID `bson:"_id"`
	// Title or Quick description for this Room
	Title string `bson:"title" json:"title"`
}

func (Room) IsIdentifiable() {}
