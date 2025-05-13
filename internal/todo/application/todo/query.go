package todo

import (
	"fish-sample/internal/todo/domain/todo"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TodoListQuery 表示查询 Todo 列表的参数
type TodoQuery struct {
	// 这里可以添加其他查询参数
	ID string `uri:"id" binding:"required,uuid"`
}

type TodoQueryHandler struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewTodoQueryHandler(db *gorm.DB, log *zap.Logger) *TodoQueryHandler {
	return &TodoQueryHandler{
		db:  db,
		log: log,
	}
}

func (h *TodoQueryHandler) Handle(query TodoQuery) (*TodoDTO, error) {
	var todoEntity todo.Todo

	// 预加载 Tasks，并按 ID 倒序排列
	if err := h.db.
		Preload("Tasks", func(db *gorm.DB) *gorm.DB {
			return db.Order("id DESC")
		}).
		First(&todoEntity, "id = ?", query.ID).
		Error; err != nil {
		h.log.Error("failed to query todo", zap.Error(err))
		return nil, err
	}

	// 转换为 DTO
	todoDTO := &TodoDTO{
		ID:    todoEntity.ID,
		Title: todoEntity.Title,
		Tasks: make([]TaskDTO, len(todoEntity.Tasks)),
	}

	for i, task := range todoEntity.Tasks {
		todoDTO.Tasks[i] = TaskDTO{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
			Completed:   task.Completed,
			TodoID:      task.TodoID,
		}
	}

	return todoDTO, nil
}
