package admin

import "starter_kit/util"

func Index(ctx util.ApiContext) error {
	return ctx.Success("admin is ok")
}