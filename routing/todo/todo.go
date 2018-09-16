package todo

import (
	"github.com/kataras/iris"
	"starter_kit/handler/todo"
	"starter_kit/middleware"
	"starter_kit/util"
)

func GetRouting(group iris.Party) {
	todoParty := group.Party("/todo")
	{
		todoParty.Get("/", middleware.ParsePaging, util.ApiHandlerWrap(todo.List))
		todoParty.Post("/", util.ApiHandlerWrap(todo.Create))
		todoParty.Get("/{id:int}", middleware.ParserPrimaryKey("id"), util.ApiHandlerWrap(todo.Detail))
		todoParty.Patch("/{id:int}", util.ApiHandlerWrap(todo.Update))
		todoParty.Delete("/{id:int}", util.ApiHandlerWrap(todo.Remove))
	}
}
