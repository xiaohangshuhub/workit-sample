package application

import (
	todo "newb-sample/internal/todo/application/todo"

	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(todo.NewCreateTodoCommandHandler),
	}

}
