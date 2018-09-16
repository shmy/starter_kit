package todo

import (
	"starter_kit/model"
	todo2 "starter_kit/model/todo"
	"starter_kit/util"
)

type Todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}

func Create(ctx util.ApiContext) error {
	body := Todo{}
	err := ctx.ReadJSON(&body)
	if err != nil {
		return ctx.Fail(err)
	}
	todo := todo2.Todo{
		Title:   body.Title,
		Content: body.Content,
	}
	err = model.GetDB().Create(&todo).Error
	if err != nil {
		return ctx.Fail(err)
	}
	return ctx.Success(todo)
}
