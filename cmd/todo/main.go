package main

import (
	"fmt"

	"mfish-sample/internal/todo/application"
	"mfish-sample/internal/todo/domain"
	"mfish-sample/internal/todo/webapi"

	_ "mfish-sample/api/todo/docs" // swagger 一定要有这行

	"github.com/gin-gonic/gin"
	"github.com/lxhanghub/go-mfish/pkg/database"
	"github.com/lxhanghub/go-mfish/pkg/host"
	"go.uber.org/zap"
)

func main() {

	// 创建服务主机构建器
	builder := host.NewWebHostBuilder()

	// 配置应用配置,内置环境变量读取和命令行参数读取
	builder.ConfigureAppConfiguration(func(build host.ConfigBuilder) {
		build.AddYamlFile("./config.yaml")
	})

	// 配置依赖注入
	builder.ConfigureServices(database.MysqlModule())

	// 领域层注入
	builder.ConfigureServices(domain.DependencyInjection()...)

	builder.ConfigureServices(application.DependencyInjection()...)

	//构建应用
	app, err := builder.Build()

	if err != nil {
		fmt.Printf("Failed to build application: %v\n", err)
		return
	}

	if app.Env.IsDevelopment {
		app.UseSwagger()
	}

	//配置请求中间件,支持跳过
	//app.UseMiddleware(middleware.NewAuthorizationMiddleware([]string{"/hello"}))

	app.UseCORS()
	// 配置路由
	app.MapRoutes(webapi.RegisterTodoRoutes)

	app.MapRoutes(func(router *gin.Engine) {
		router.GET("/ping", func(c *gin.Context) {

			c.JSON(200, gin.H{"message": "hello world"})
		})
	})

	// 运行应用
	if err := app.Run(); err != nil {
		app.Logger().Error("Error running application", zap.Error(err))
	}
}
