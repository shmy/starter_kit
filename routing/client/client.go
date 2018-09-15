package client

import (
	"github.com/kataras/iris"
	"starter_kit/handler/client"
	"starter_kit/util"
)

func GetRouting(group iris.Party) {
	group.Get("/", util.ApiHandlerWrap(client.Index))
}
