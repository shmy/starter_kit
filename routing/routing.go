package routing

import (
	"github.com/kataras/iris"
	"starter_kit/routing/admin"
	"starter_kit/routing/client"
	"starter_kit/routing/svcd"
)

func RouteMapping(app *iris.Application) {
	// 静态服务
	app.StaticWeb("/static", "./public")
	svcdGroup := app.Party("/svcd")
	clientGroup := app.Party("/client")
	adminGroup := app.Party("/admin")
	svcd.GetRouting(svcdGroup)
	client.GetRouting(clientGroup)
	admin.GetRouting(adminGroup)
}
