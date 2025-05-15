package webapi

import (
	"mfish-sample/internal/todo/application/todo"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func RegisterTodoRoutes(
	router *gin.Engine, //gin 放在第一位
	log *zap.Logger, // 日志
	create *todo.CreateTodoCommandHandler, // 创建
	todoList *todo.TodoListQueryHandler, //列表
	addTask *todo.AddTodoTaskCommandHandler, //添加任务
	todoQuery *todo.TodoQueryHandler, //查询
	markAsCompleted *todo.MarkAsCompletedCommandHandler, //查询
) {

	// 创建路由组
	group := router.Group("/todos")

	// 创建路由
	group.POST("", CreateTodoHandler(create, log))
	group.GET("", TodoListQueryHandler(todoList, log))
	group.POST("/task", AddTodoTaskHandler(addTask, log))
	group.GET("/:id", TodoQueryHandler(todoQuery, log))
	group.POST("/completed", MarkAsCompletedHandler(markAsCompleted, log))
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

// AddTodoTaskHandler godoc
// @Summary 添加任务
// @Description 为指定的待办事项添加任务
// @Tags Todos
// @Accept json
// @Produce json
// @Param data body todo.AddTodoTaskCommand true "请求参数"
// @Success 200 {object} Response[bool]
// @Failure 400 {object} Response[any]
// @Failure 500 {object} Response[any]
// @Router /todos/task [post]
func AddTodoTaskHandler(handler *todo.AddTodoTaskCommandHandler, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmd todo.AddTodoTaskCommand

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Error("params error", zap.Error(err))
			Fail(c, 400, "参数错误: "+err.Error())
			return
		}

		result, err := handler.Handle(cmd)
		if err != nil {
			log.Error("add task error", zap.Error(err))
			Fail(c, 500, "添加任务失败: "+err.Error())
			return
		}
		Success(c, result)
	}
}

// TodoQueryHandler godoc
// @Summary 查询Todo
// @Description 查询指定ID的待办事项
// @Tags Todos
// @Accept json
// @Produce json
// @Param id path string true "待办事项ID"
// @Success 200 {object} Response[todo.TodoDTO]
// @Failure 400 {object} Response[any]
// @Failure 500 {object} Response[any]
// @Router /todos/{id} [get]
func TodoQueryHandler(handler *todo.TodoQueryHandler, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {

		var query todo.TodoQuery

		if err := c.ShouldBindUri(&query); err != nil {
			log.Error("uri bind error", zap.Error(err))
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

// MarkAsCompletedHandler godoc
// @Summary 标记任务为完成
// @Description 将指定的任务标记为完成
// @Tags Todos
// @Accept json
// @Produce json
// @Param data body todo.MarkAsCompletedCommand true "请求参数"
// @Success 200 {object} Response[bool]
// @Failure 400 {object} Response[any]
// @Failure 500 {object} Response[any]
// @Router /todos/completed [post]
func MarkAsCompletedHandler(handler *todo.MarkAsCompletedCommandHandler, log *zap.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmd todo.MarkAsCompletedCommand

		if err := c.ShouldBindJSON(&cmd); err != nil {
			log.Error("params error", zap.Error(err))
			Fail(c, 400, "参数错误: "+err.Error())
			return
		}

		result, err := handler.Handle(cmd)
		if err != nil {
			log.Error("mark as completed error", zap.Error(err))
			Fail(c, 500, "标记完成失败: "+err.Error())
			return
		}
		Success(c, result)
	}
}
