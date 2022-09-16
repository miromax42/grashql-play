package array

import (
	"github.com/google/uuid"

	"github.com/miromax42/grashql-play/graph/model"
	"github.com/miromax42/grashql-play/utils"
)

type User struct {
	users []*model.User
}

func (a *User) GetUser(id uuid.UUID) (*model.User, error) {
	for i := range a.users {
		if a.users[i].ID == id {
			return a.users[i], nil
		}
	}

	return nil, utils.ErrNotExists
}

func (a *User) CreateUser(user model.NewUser) *model.User {
	u := &model.User{
		ID:   uuid.New(),
		Name: user.Name,
	}
	a.users = append(a.users, u)

	return a.users[len(a.users)-1]
}

func (a *User) ListUser() []*model.User {
	return a.users
}
