package todo

import (
	"newb-sample/internal/todo/domain/todo"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type AddTodoTaskCommand struct {
	TodoID      uuid.UUID `json:"todoId" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
	Title       string    `json:"title" example:"Buy milk"`
	Description *string   `json:"description" example:"From supermarket"`
}

type AddTodoTaskCommandHandler struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewAddTodoTaskCommandHandler(db *gorm.DB, log *zap.Logger) *AddTodoTaskCommandHandler {
	return &AddTodoTaskCommandHandler{
		db:  db,
		log: log,
	}
}

func (h *AddTodoTaskCommandHandler) Handle(cmd AddTodoTaskCommand) (bool, error) {

	todo := todo.Todo{}

	// 使用 Preload 加载关联的 Tasks
	result := h.db.Preload("Tasks").First(&todo, "id = ?", cmd.TodoID)

	if result.Error != nil {
		h.log.Error("failed to query todoList", zap.Error(result.Error))
		return false, result.Error
	}

	err := todo.AddTask(uuid.New(), cmd.Title, cmd.Description)

	if err != nil {
		h.log.Error("failed to add task", zap.Error(err))
		return false, err
	}

	tx := h.db.Save(todo)

	if tx.Error != nil {
		h.log.Error("failed to save task", zap.Error(tx.Error))
		return false, tx.Error
	}

	return true, nil
}
