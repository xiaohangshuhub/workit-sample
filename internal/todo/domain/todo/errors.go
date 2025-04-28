package todo

type TodoError struct {
	Message string
}

func (e TodoError) Error() string {
	return e.Message
}

var (
	ErrEmptyTodoTitle    = TodoError{Message: "待办事项标题不能为空"}
	ErrTodoAlreadyExists = TodoError{Message: "待办事项已存在"}
	ErrEmptyTaskTitle    = TodoError{Message: "任务标题不能为空"}
	ErrTaskNotFound      = TodoError{Message: "任务未找到"}
)
