package api

import (
	todoapp "newb-sample/internal/todo/application/todo-app"

	"github.com/gin-gonic/gin"
)

func RegisterTodoRoutes(
	router *gin.Engine, //gin 必须放在第一位,框架要求否则会报错
	create *todoapp.CreateTodoCommandHandler) {
	router.GET("/todos", CreateTodoHandler(create))
}

// CreateTodoHandler godoc
// @Summary 创建新的Todo
// @Description 创建新的Todo
// @Tags Todos
// @Accept json
// @Produce json
// @Param data body todoapp.CreateTodoCommand true "Todo请求参数"
// @Success 200 {object} todoapp.CreateTodoResult
// @Failure 400 {object} gin.H
// @Failure 500 {object} gin.H
// @Router /todos [post]
func CreateTodoHandler(handler *todoapp.CreateTodoCommandHandler) gin.HandlerFunc {
	return func(c *gin.Context) {
		var cmd todoapp.CreateTodoCommand
		if err := c.ShouldBindJSON(&cmd); err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}

		result, err := handler.Handle(cmd)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}

		c.JSON(200, result)
	}
}
