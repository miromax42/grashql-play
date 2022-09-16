package array

import (
	"github.com/google/uuid"

	"github.com/miromax42/grashql-play/graph/model"
	"github.com/miromax42/grashql-play/repo"
	"github.com/miromax42/grashql-play/utils"
)

type TODO struct {
	users repo.User
	todos []*model.Todo
}

func (a *TODO) UpdateTODO(values model.UpdateTodo) (*model.Todo, error) {
	todo, err := a.GetTODO(values.ID)
	if err != nil {
		return nil, err
	}

	if values.Done != nil {
		todo.Done = *values.Done
	}
	if values.Text != nil {
		todo.Text = *values.Text
	}

	return todo, nil
}

func NewTODO(users repo.User) repo.Todo {
	return &TODO{
		users: users,
	}
}

func (a *TODO) CreateTODO(todo model.NewTodo) (*model.Todo, error) {
	user, err := a.users.GetUser(todo.UserID)
	if err != nil {
		return nil, utils.ErrNotExists
	}
	t := &model.Todo{
		ID:   uuid.New(),
		Text: todo.Text,
		Done: false,
		User: user,
	}
	a.todos = append(a.todos, t)

	return a.todos[len(a.todos)-1], nil
}

func (a *TODO) GetTODO(id uuid.UUID) (*model.Todo, error) {
	for i := range a.todos {
		if a.todos[i].ID == id {
			return a.todos[i], nil
		}
	}

	return nil, utils.ErrNotExists
}

func (a *TODO) ListTODO() []*model.Todo {
	return a.todos
}
