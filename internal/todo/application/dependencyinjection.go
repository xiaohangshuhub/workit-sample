package application

import (
	todoapp "newb-sample/internal/todo/application/todo-app"

	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(todoapp.NewCreateTodoCommandHandler),
	}

}
