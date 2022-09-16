package array

import (
	"github.com/google/uuid"

	"gql-play/graph/model"
	"gql-play/repo"
)

type TODO struct {
	users repo.User
	todos []*model.Todo
}

func NewTODO(users repo.User) repo.Todo {
	return &TODO{
		users: users,
	}
}

func (a *TODO) CreateTODO(todo model.NewTodo) *model.Todo {
	t := &model.Todo{
		ID:   uuid.NewString(),
		Text: todo.Text,
		Done: false,
		User: a.users.GetUser(uuid.Must(uuid.Parse(todo.UserID))),
	}
	a.todos = append(a.todos, t)

	return a.todos[len(a.todos)-1]
}

func (a *TODO) GetTODO(id uuid.UUID) *model.Todo {
	for i := range a.todos {
		if a.todos[i].ID == id.String() {
			return a.todos[i]
		}
	}

	return nil
}

func (a *TODO) ListTODO() []*model.Todo {
	return a.todos
}
