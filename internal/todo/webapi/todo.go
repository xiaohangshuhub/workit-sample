package webapi

import (
	"newb-sample/internal/todo/application/todo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterTodoRoutes(
	router *gin.Engine, //gin 放在第一位
	log *zap.Logger,
	create *todo.CreateTodoCommandHandler,
	todoList *todo.TodoListQueryHandler) {

	// 创建路由组
	group := router.Group("/todos")

	// 创建路由
	group.POST("", CreateTodoHandler(create, log))
	group.GET("", TodoListQueryHandler(todoList, log))
}

// CreateTodoHandler godoc
// @Summary 创建Todo
// @Description 创建新的待办事项
// @Tags Todos
// @Accept json
// @Produce json
// @Param data body todo.CreateTodoCommand true "请求参数"
// @Success 200 {object} Response[todo.CreateTodoResult]
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

// TodoListQueryHandler godoc
// @Summary 查询Todo列表
// @Description 查询所有匹配条件的待办事项
// @Tags Todos
// @Accept json
// @Produce json
// @Param title query string false "任务标题"
// @Param page query int false "页码"
// @Param size query int false "每页大小"
// @Success 200 {object} Response[[]todo.TodoDTO]
// @Failure 400 {object} Response[any]
// @Failure 500 {object} Response[any]
// @Router /todos [get]
func TodoListQueryHandler(handler *todo.TodoListQueryHandler, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var query todo.TodoListQuery

		if err := c.ShouldBindQuery(&query); err != nil {
			log.Error("params error", zap.Error(err))
			Fail(c, 400, "参数错误: "+err.Error())
			return
		}

		result, err := handler.Handle(query)
		if err != nil {
			log.Error("query error", zap.Error(err))
			Fail(c, 500, "查询失败: "+err.Error())
			return
		}
		Success(c, result)
	}
}
