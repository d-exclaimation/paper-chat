// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

type Identifiable interface {
	IsIdentifiable()
}

type Room struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

func (Room) IsIdentifiable() {}
