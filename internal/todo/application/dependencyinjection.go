package application

import (
	todo "workit-sample/internal/todo/application/todo"

	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(todo.NewCreateTodoCommandHandler),
		fx.Provide(todo.NewTodoListQueryHandler),
		fx.Provide(todo.NewAddTodoTaskCommandHandler),
		fx.Provide(todo.NewTodoQueryHandler),
		fx.Provide(todo.NewMarkAsCompletedCommandHandler),
	}

}
