package repo

import (
	"github.com/google/uuid"

	"gql-play/graph/model"
)

type User interface {
	CreateUser(user model.NewUser) *model.User
	GetUser(id uuid.UUID) *model.User
	ListUser() []*model.User
}
