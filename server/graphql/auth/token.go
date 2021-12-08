//
//  token.go
//  auth
//
//  Created by d-exclaimation on 11:35 AM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package auth

import (
	"context"
	"github.com/d-exclaimation/paper-chat/db/users"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"strings"
)

const Key = "ContextAuthKey"

func Auth(db *mongo.Database, ctx context.Context) <-chan *model.User {
	task := make(chan *model.User)

	go func() {
		token := GetAccessToken(ctx)
		if token == nil {
			task <- nil
			close(task)
			return
		}

		id, err := Unsign(*token)

		if err != nil {
			task <- nil
			close(task)
			return
		}

		oid, err := primitive.ObjectIDFromHex(id)

		if err != nil {
			task <- nil
			close(task)
			return
		}

		user := <-users.GetById(db, oid, ctx)
		task <- user
		close(task)
	}()

	return task
}

func GetAccessToken(ctx context.Context) *string {
	key, ok := ctx.Value(Key).(string)
	if !ok {
		return nil
	}
	return &key
}

func Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
		header := req.Header.Get("Authorization")
		bearerToken := strings.Split(header, " ")

		if len(bearerToken) < 2 || strings.ToLower(bearerToken[0]) != "bearer" {
			next.ServeHTTP(res, req)
			return
		}

		token := bearerToken[1]

		ctx := context.WithValue(req.Context(), Key, token)

		req = req.WithContext(ctx)
		next.ServeHTTP(res, req)
	})
}
