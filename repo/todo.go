package repo

import (
	"github.com/google/uuid"

	"github.com/miromax42/grashql-play/graph/model"
)

type Todo interface {
	CreateTODO(todo model.NewTodo) *model.Todo
	GetTODO(id uuid.UUID) *model.Todo
	ListTODO() []*model.Todo
}
