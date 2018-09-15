package admin

import (
	"starter_kit/model"
	"starter_kit/model/menu"
	"starter_kit/util"
)

func Index(ctx util.ApiContext) error {
	menus := make([]menu.Menu, 0)
	err := model.GetDB().Select("name").Find(&menus).Error
	if err != nil {
		return ctx.Fail(err, 200)
	}
	return ctx.Success(menus)
}
