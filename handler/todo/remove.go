package todo

import (
	"errors"
	"github.com/jinzhu/gorm"
	"starter_kit/model"
	todo2 "starter_kit/model/todo"
	"starter_kit/util"
)

func Remove(ctx util.ApiContext) error {
	id, err := ctx.Params().GetInt("id")
	if err != nil {
		return ctx.Fail(errors.New("id格式不正确"))
	}
	todo := todo2.Todo{}
	err = model.GetDB().First(&todo, id).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return ctx.Fail(errors.New("记录不存在"))
		}
		return ctx.Fail(err)
	}
	err = model.GetDB().Delete(&todo).Error
	if err != nil {
		return ctx.Fail(err)
	}
	return ctx.Success(todo)
}
