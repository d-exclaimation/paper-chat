//
// main.go
// paper-chat
//
// Created by d-exclaimation on 00:00.
//

package main

import (
	_ "github.com/99designs/gqlgen/graphql"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	graph "github.com/d-exclaimation/paper-chat/graphql"
	"github.com/d-exclaimation/paper-chat/graphql/gql"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"time"
)

func main() {
	r := chi.NewRouter()

	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000"},
		AllowCredentials: true,
	}))

	resolver := &graph.Resolver{}

	srv := handler.New(gql.NewExecutableSchema(gql.Config{
		Resolvers:  resolver,
		Directives: gql.DirectiveRoot{},
		Complexity: gql.ComplexityRoot{},
	}))

	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	srv.Use(extension.Introspection{})

	r.Get("/playground", playground.Handler("PaperChat", "/graphql"))
	r.Handle("/graphql", srv)

	if err := http.ListenAndServe(":4000", r); err != nil {
		log.Fatalln(err)
	}
}
