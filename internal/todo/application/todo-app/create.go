package todoapp

import (
	"newb-sample/internal/todo/domain/todo"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreateTodoCommand struct {
	Title       string  `json:"title" validate:"required"`
	Description *string `json:"description"`
}

type CreateTodoResult struct {
	Sucess bool `json:"success"`
}

type CreateTodoCommandHandler struct {
	db      *gorm.DB
	log     *zap.Logger
	manager *todo.TodoManager
}

func NewCreateTodoCommandHandler(db *gorm.DB, log *zap.Logger, todoManager *todo.TodoManager) *CreateTodoCommandHandler {
	return &CreateTodoCommandHandler{
		db:      db,
		log:     log,
		manager: todoManager,
	}
}

func (h *CreateTodoCommandHandler) Handle(cmd CreateTodoCommand) (*CreateTodoResult, error) {

	todo, err := h.manager.CreateTodo(cmd.Title, cmd.Description)

	if err != nil {
		return nil, err
	}

	tx := h.db.Create(todo)

	if tx.Error != nil {
		return nil, tx.Error
	}

	return &CreateTodoResult{
		Sucess: true,
	}, nil
}
