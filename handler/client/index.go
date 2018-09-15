package client

import "starter_kit/util"

func Index(ctx util.ApiContext) error {
	return ctx.Success("client is ok")
}
