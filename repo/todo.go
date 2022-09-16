package repo

import (
	"github.com/google/uuid"

	"github.com/miromax42/grashql-play/graph/model"
)

type Todo interface {
	CreateTODO(todo model.NewTodo) (*model.Todo, error)
	GetTODO(id uuid.UUID) (*model.Todo, error)
	UpdateTODO(todo model.UpdateTodo) (*model.Todo, error)
	ListTODO() []*model.Todo
}
