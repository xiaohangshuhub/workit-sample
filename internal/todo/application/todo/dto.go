package todo

import "github.com/google/uuid"

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
