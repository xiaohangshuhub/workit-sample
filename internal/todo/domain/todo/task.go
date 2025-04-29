package todo

import (
	"newb-sample/pkg/ddd"

	"github.com/google/uuid"
)

type Task struct {
	ddd.Entity[uuid.UUID]
	Title       string    `json:"title" gorm:"column:title"`
	Description string    `json:"description" gorm:"column:description"`
	Completed   bool      `json:"completed" gorm:"column:completed"`
	TodoID      uuid.UUID `json:"todo_id" gorm:"column:todo_id"`
}
