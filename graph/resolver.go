package graph

import (
	"gql-play/repo"
	"gql-play/repo/array"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	userRepo repo.User
	todoRepo repo.Todo
}

func NewResolver() *Resolver {
	userRepo := &array.User{}
	return &Resolver{
		userRepo: userRepo,
		todoRepo: array.NewTODO(userRepo),
	}
}
