//
//  constant.go.go
//  config
//
//  Created by d-exclaimation on 8:50 PM.
//  Copyright © 2021 d-exclaimation. All rights reserved.
//

package config

import (
	"os"
	"strings"
)

func Port() string {
	if port, ok := os.LookupEnv("PORT"); ok {
		return port
	}
	return "4000"
}

func Runtime() string {
	var (
		options   = []string{"production", "development", "maintenance", "testing"}
		isRuntime = func(r string) bool {
			for _, runtime := range options {
				if r == runtime {
					return true
				}
			}
			return false
		}
	)
	if runtime, ok := os.LookupEnv("GO_ENV"); ok && isRuntime(runtime) {
		return runtime
	}
	return options[1]
}

func IsProd() bool {
	return strings.ToLower(Runtime()) == "production"
}

func MongoURI() string {
	if uri, ok := os.LookupEnv("MONGO_URI"); ok {
		return uri
	}
	return "mongodb://localhost/paper-chat"
}
