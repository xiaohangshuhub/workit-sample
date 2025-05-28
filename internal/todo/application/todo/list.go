package todo

import (
	"workit-sample/internal/todo/domain/todo"

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

func NewTodoListQueryHandler(db *gorm.DB, log *zap.Logger) *TodoListQueryHandler {
	return &TodoListQueryHandler{
		db:  db,
		log: log,
	}
}

func (h *TodoListQueryHandler) Handle(query TodoListQuery) ([]TodoDTO, error) {
	var todos []todo.Todo

	// 查询所有待办事项，按 ID 倒序排列
	if err := h.db.Order("id DESC").Find(&todos).Error; err != nil {
		h.log.Error("failed to query todo list", zap.Error(err))
		return nil, err
	}

	todoDTOs := make([]TodoDTO, len(todos))

	for i, t := range todos {
		todoDTOs[i] = TodoDTO{
			ID:          t.ID,
			Title:       t.Title,
			Description: t.Description,
			Completed:   t.Completed,
			Tasks:       []TaskDTO{}, // 为空但保持字段一致性
		}
	}

	return todoDTOs, nil
}
