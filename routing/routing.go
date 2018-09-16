package routing

import (
	"github.com/kataras/iris"
	"starter_kit/routing/svcd"
	"starter_kit/routing/todo"
)

func RouteMapping(app *iris.Application) {
	// 静态服务
	app.StaticWeb("/static", "./public")
	// 系统监控
	svcdGroup := app.Party("/svcd")
	svcd.GetRouting(svcdGroup)

	// api v1
	apiV1Party := app.Party("/api/v1")
	{
		todo.GetRouting(apiV1Party)
	}
}
