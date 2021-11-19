//
//  user.go
//  model
//
//  Created by d-exclaimation on 7:46 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

import "go.mongodb.org/mongo-driver/bson/primitive"

// User account for PaperChat
type User struct {
	// User ID
	OID primitive.ObjectID `bson:"_id"`
	// Unique username for this User
	Username string `bson:"username" json:"username"`
}

func (User) IsIdentifiable() {}
