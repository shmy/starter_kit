package todo

import (
	"fmt"
	"starter_kit/middleware"
	"starter_kit/model"
	"starter_kit/model/todo"
	"starter_kit/util"
)

func List(ctx util.ApiContext) error {
	paging := ctx.Values().Get("paging").(middleware.Paging)
	fmt.Println("分页参数", paging)
	todos := make([]todo.Todo, 0)
	err := model.GetDB().Select("id, title, content, created_at").Find(&todos).Error
	if err != nil {
		return ctx.Fail(err)
	}
	return ctx.Success(todos)
}
