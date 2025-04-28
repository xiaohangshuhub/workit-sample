package domain

import (
	"newb-sample/internal/todo/domain/todo"

	"go.uber.org/fx"
)

func DependencyInjection() []fx.Option {

	return []fx.Option{
		fx.Provide(todo.NewTodoManager),
	}

}
