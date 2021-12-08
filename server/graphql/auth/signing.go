//
//  signing.go
//  auth
//
//  Created by d-exclaimation on 12:16 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package auth

import (
	"errors"
	"fmt"
	"github.com/d-exclaimation/paper-chat/config"
	"github.com/d-exclaimation/paper-chat/graphql/model"
	"github.com/golang-jwt/jwt"
	"time"
)

func Sign(user *model.User) (string, time.Time, error) {
	dur, _ := time.ParseDuration("720h")
	exp := time.Now().Add(dur).UTC()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":  user.OID.Hex(),
		"exp": exp.Unix(),
	})

	str, err := token.SignedString(config.TokenSecret())
	return str, exp, err
}

func Unsign(token string) (string, error) {
	parsed, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return config.TokenSecret(), nil
	})

	if claims, ok := parsed.Claims.(jwt.MapClaims); ok && parsed.Valid {
		if id, ok2 := claims["id"].(string); ok2 {
			return id, nil
		}
		return "", errors.New("no id found")
	} else {
		return "", err
	}
}
