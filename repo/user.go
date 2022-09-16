package repo

import (
	"github.com/google/uuid"

	"github.com/miromax42/grashql-play/graph/model"
)

type User interface {
	CreateUser(user model.NewUser) *model.User
	GetUser(id uuid.UUID) (*model.User, error)
	ListUser() []*model.User
}
