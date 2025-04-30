package todo

import (
	"newb-sample/internal/todo/domain/todo"

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

	todo := todo.Todo{}

	// 使用 Preload 加载关联的 Tasks
	result := h.db.Preload("Tasks").First(&todo, "id = ?", query.ID)

	if result.Error != nil {
		h.log.Error("failed to query todoList", zap.Error(result.Error))
		return nil, result.Error
	}

	// 将 todo 转换为 DTO
	tasksDTO := make([]TaskDTO, len(todo.Tasks))
	for i, task := range todo.Tasks {
		tasksDTO[i] = TaskDTO{
			ID:          task.ID,
			Title:       task.Title,
			Description: task.Description,
		}
	}

	todoDTO := &TodoDTO{
		ID:    todo.ID,
		Title: todo.Title,
		Tasks: tasksDTO,
	}

	return todoDTO, nil
}
