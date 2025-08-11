package main

import (
	"fmt"

	"workit-sample/internal/todo/application"
	"workit-sample/internal/todo/domain"
	"workit-sample/internal/todo/webapi"

	_ "workit-sample/api/todo/docs" // swagger 一定要有这行

	"github.com/gin-contrib/cors"
	"github.com/xiaohangshuhub/go-workit/pkg/database"
	"github.com/xiaohangshuhub/go-workit/pkg/workit"
	"go.uber.org/zap"
)

func main() {

	// 创建服务主机构建器
	builder := workit.NewWebAppBuilder()

	// 配置应用配置,内置环境变量读取和命令行参数读取
	builder.AddConfig(func(build workit.ConfigBuilder) {
		build.AddYamlFile("./config.yaml")
	})

	// 配置依赖注入
	builder.AddServices(database.MysqlModule())

	// 领域层注入
	builder.AddServices(domain.DependencyInjection()...)

	builder.AddServices(application.DependencyInjection()...)

	builder.AddAuthentication().AddJwtBearer(func(options *workit.JwtBearerOptions) {

		options.TokenValidationParameters = workit.TokenValidationParameters{
			ValidateIssuer:           true,
			ValidateAudience:         true,
			ValidateLifetime:         true,
			ValidateIssuerSigningKey: true,
			SigningKey:               []byte("secret"),
			ValidIssuer:              "sample",
			ValidAudience:            "sample",
			RequireExpiration:        true,
		}
	})

	builder.AddAuthorization(workit.AuthorizeOptions{
		Routes: []workit.Route{
			{
				Path: "/test/{action}",
				Methods: []workit.RequestMethod{
					workit.GET,
				},
			},
		},

		Policies: []string{"admin_role_policy"},
	}).
		RequireRole("admin_role_policy", "Admin")

	//构建应用
	app, err := builder.Build()

	if err != nil {
		fmt.Printf("Failed to build application: %v\n", err)
		return
	}

	if app.Env().IsDevelopment {
		app.UseSwagger()
	}

	// 配置跨域
	app.UseCORS(func(c *cors.Config) {
		c.AllowAllOrigins = true
		c.AllowCredentials = true
	})

	// 配置鉴权
	app.UseAuthentication()

	// 配置授权
	app.UseAuthorization()

	// 配置路由
	app.MapRoutes(webapi.RegisterTodoRoutes)

	// 运行应用
	if err := app.Run(); err != nil {
		app.Logger().Error("Error running application", zap.Error(err))
	}
}
