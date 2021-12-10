//
//  message.go
//  model
//
//  Created by d-exclaimation on 11:12 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package model

type Message struct {
	Value     string `bson:"value"`
	CreatedAt string `bson:"created_at"`
}
