package todo

import (
	"errors"

	"github.com/google/uuid"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type TodoManager struct {
	db  *gorm.DB
	log *zap.Logger
}

func NewTodoManager(db *gorm.DB, log *zap.Logger) (*TodoManager, error) {
	return &TodoManager{
		db:  db,
		log: log,
	}, nil
}
func (m *TodoManager) CreateTodo(title string, desc *string) (*Todo, error) {

	// 检查标题是否存在
	var todo *Todo

	tx := m.db.Where("title = ?", title).First(&todo)

	if tx.Error != nil && !errors.Is(tx.Error, gorm.ErrRecordNotFound) {
		m.log.Error("failed to check todo title", zap.Error(tx.Error))
		return nil, tx.Error
	}

	if tx.RowsAffected != 0 {
		m.log.Error("todo already exists", zap.String("title", title))
		return nil, ErrTodoAlreadyExists
	}

	todo, err := NewTodo(uuid.New(), title)

	todo.Description = desc

	if err != nil {
		m.log.Error("failed to create todo", zap.Error(err))
		return nil, err
	}

	return todo, nil
}
