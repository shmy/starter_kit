package todo

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/core/errors"
	"starter_kit/model"
	todo2 "starter_kit/model/todo"
	"starter_kit/util"
)

func Update(ctx util.ApiContext) error {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		return ctx.Fail(errors.New("id格式不正确"))
	}
	body := todo2.Todo{}
	err = ctx.ReadJSON(&body)
	if err != nil {
		return ctx.Fail(err)
	}
	todo := todo2.Todo{}
	err = model.GetDB().First(&todo, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Fail(errors.New("记录不存在"))
		}
		return ctx.Fail(err)
	}
	err = model.GetDB().Model(&todo).Updates(map[string]string{
		"title":   body.Title,
		"content": body.Content,
	}).Error
	if err != nil {
		return ctx.Fail(err)
	}
	return ctx.Success(todo)
}
