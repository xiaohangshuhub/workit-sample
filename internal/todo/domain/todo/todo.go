package todo

import (
	"newb-sample/pkg/ddd"
	"newb-sample/pkg/tools/str"

	"github.com/google/uuid"
)

type Todo struct {
	ddd.AggregateRoot[uuid.UUID]
	Title       string  `json:"title" gorm:"column:title"`
	Description *string `json:"description" gorm:"column:description"`
	Completed   bool    `json:"completed" gorm:"column:completed"`
	Tasks       []Task  `json:"tasks" gorm:"foreignKey:TodoID;references:ID"`
}

func NewTodo(id uuid.UUID, title string) (*Todo, error) {
	if str.IsEmptyOrWhiteSpace(title) {
		return nil, ErrEmptyTodoTitle
	}
	return &Todo{
		AggregateRoot: ddd.NewAggregateRoot(id),
		Title:         title,
		Completed:     false,
	}, nil
}

func (t *Todo) AddTask(taskId uuid.UUID, title string, description *string) error {
	if str.IsEmptyOrWhiteSpace(title) {
		return ErrEmptyTaskTitle
	}

	task := Task{
		Entity:      ddd.NewEntity(taskId),
		Title:       title,
		Description: *description,
		Completed:   false,
	}

	t.Tasks = append(t.Tasks, task)
	return nil
}

func (t *Todo) UpdateTitle(title string) error {
	if str.IsEmptyOrWhiteSpace(title) {
		return ErrEmptyTodoTitle
	}
	t.Title = title
	return nil
}

func (t *Todo) RemoveTask(taskId uuid.UUID) error {
	for i, task := range t.Tasks {
		if task.ID == taskId {
			t.Tasks = append(t.Tasks[:i], t.Tasks[i+1:]...)
			return nil
		}
	}
	return ErrTaskNotFound
}

func (t *Todo) RemoveTasks(taskIds []uuid.UUID) error {
	for _, taskId := range taskIds {
		if err := t.RemoveTask(taskId); err != nil {
			return err
		}
	}
	return nil
}
