package todo

import (
	"starter_kit/model"
	todo2 "starter_kit/model/todo"
	"starter_kit/util"
)

func Create(ctx util.ApiContext) error {
	body := todo2.Todo{}
	err := ctx.ReadJSON(&body)
	if err != nil {
		return ctx.Fail(err)
	}
	err = model.GetDB().Create(&body).Error
	if err != nil {
		return ctx.Fail(err)
	}
	return ctx.Success(body)
}
