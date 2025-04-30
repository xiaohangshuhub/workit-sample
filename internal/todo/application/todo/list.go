package todo

import (
	"newb-sample/internal/todo/domain/todo"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

// TodoListQuery 表示查询 Todo 列表的参数
type TodoListQuery struct {
	// 这里可以添加其他查询参数
	Title string `form:"title" example:"Buy milk"` // 可选标题关键词
	Page  int    `form:"page" example:"1"`         // 页码
	Size  int    `form:"size" example:"10"`        // 每页条数
}

type TodoListQueryHandler struct {
	db  *gorm.DB
	log *zap.Logger
}

// TodoItemDTO 是用于 Swagger 展示的简化结构
type TodoDTO struct {
	ID          uuid.UUID `json:"id" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
	Title       string    `json:"title" example:"Buy milk"`
	Description *string   `json:"description" example:"From supermarket"`
	Completed   bool      `json:"completed" example:"false"`
	Tasks       []TaskDTO `json:"tasks"`
}

type TaskDTO struct {
	ID          uuid.UUID `json:"id" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
	TodoID      uuid.UUID `json:"todoId" example:"b19e6f4c-3d51-4f7e-9a6e-f32d28a3f111"`
	Title       string    `json:"title" example:"Buy milk"`
	Description *string   `json:"description" example:"From supermarket"`
	Completed   bool      `json:"completed" example:"false"`
}

func NewTodoListQueryHandler(db *gorm.DB, log *zap.Logger) *TodoListQueryHandler {
	return &TodoListQueryHandler{
		db:  db,
		log: log,
	}
}

func (h *TodoListQueryHandler) Handle(query TodoListQuery) ([]TodoDTO, error) {

	// 查询所有的待办事项
	todoList := []todo.Todo{}

	// query  我这里没有使用条件查询

	result := h.db.Find(&todoList)

	if result.Error != nil {
		h.log.Error("failed to query todoList", zap.Error(result.Error))
		return nil, result.Error
	}

	todoDTOList := make([]TodoDTO, len(todoList))

	// 将 todoList 转换为 TodoDTO
	for i, t := range todoList {
		todoDTOList[i] = TodoDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			Tasks:       make([]TaskDTO, 0),
		}
	}

	return todoDTOList, nil
}
