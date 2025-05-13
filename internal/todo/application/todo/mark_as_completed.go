package todo

import (
	"fish-sample/internal/todo/domain/todo"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type MarkAsCompletedCommand struct {
	TodoID uuid.UUID `json:"todoId" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
	TaskID uuid.UUID `json:"taskId" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
}

type MarkAsCompletedCommandHandler struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewMarkAsCompletedCommandHandler(db *gorm.DB, log *zap.Logger) *MarkAsCompletedCommandHandler {
	return &MarkAsCompletedCommandHandler{
		db:  db,
		log: log,
	}
}

func (h *MarkAsCompletedCommandHandler) Handle(cmd MarkAsCompletedCommand) (bool, error) {

	todo := todo.Todo{}

	// 使用 Preload 加载关联的 Tasks
	result := h.db.Preload("Tasks").First(&todo, "id = ?", cmd.TodoID)

	if result.Error != nil {
		h.log.Error("failed to query todoList", zap.Error(result.Error))
		return false, result.Error
	}

	err := todo.MarkAsCompleted(cmd.TaskID)

	if err != nil {
		h.log.Error("failed to add task", zap.Error(err))
		return false, err
	}

	tx := h.db.Session(&gorm.Session{FullSaveAssociations: true}).Save(&todo)

	if tx.Error != nil {
		h.log.Error("failed to save task", zap.Error(tx.Error))
		return false, tx.Error
	}

	return true, nil
}
