package graph

import (
	"github.com/miromax42/grashql-play/repo"
	"github.com/miromax42/grashql-play/repo/array"
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
