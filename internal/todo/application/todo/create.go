package todo

import (
	"newb-sample/internal/todo/domain/todo"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

type CreateTodoCommand struct {
	Title       string  `json:"title" validate:"required"` // 标题
	Description *string `json:"description"`               // 描述
}

type CreateTodoResult struct {
	Sucess bool `json:"success"` // 是否成功
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
		h.log.Error("failed to create todo", zap.Error(err))
		return nil, err
	}

	tx := h.db.Create(todo)

	if tx.Error != nil {
		h.log.Error("failed to save todo", zap.Error(tx.Error))
		return nil, tx.Error
	}

	return &CreateTodoResult{
		Sucess: true,
	}, nil
}
