package webapi

import (
	"newb-sample/internal/todo/application/todo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterTodoRoutes(
	router *gin.Engine, //gin 放在第一位
	log *zap.Logger,
	create *todo.CreateTodoCommandHandler) {
	router.POST("/todos", CreateTodoHandler(create, log))
}

// CreateTodoHandler godoc
// @Summary 创建Todo
// @Description 创建新的待办事项
// @Tags Todos
// @Accept json
// @Produce json
// @Param data body todoapp.CreateTodoCommand true "请求参数"
// @Success 200 {object} Response[todoapp.CreateTodoResult]
// @Failure 400 {object} Response[any]
// @Failure 500 {object} Response[any]
// @Router /todos [post]
func CreateTodoHandler(handler *todo.CreateTodoCommandHandler, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var cmd todo.CreateTodoCommand

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Error("params error", zap.Error(err))
			Fail(c, 400, "参数错误: "+err.Error())
			return
		}

		result, err := handler.Handle(cmd)

		if err != nil {
			log.Error("create error", zap.Error(err))
			Fail(c, 500, "创建失败: "+err.Error())
			return
		}
		Success(c, result)
	}
}
