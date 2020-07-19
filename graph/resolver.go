package graph

import "github.com/jpastorm/gqlgen-todos/graph/model"

// This file will not be regenerated automatically.
//go:generate go run github.com/99designs/gqlgen
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	users []*model.User
}
