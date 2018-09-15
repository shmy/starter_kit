package admin

import (
	"github.com/kataras/iris"
	"starter_kit/handler/admin"
	"starter_kit/util"
)

func GetRouting(group iris.Party) {
	group.Get("/", util.ApiHandlerWrap(admin.Index))
}
