package util

import (
	"github.com/kataras/iris"
	"github.com/kataras/iris/context"
)

type ApiContext struct {
	iris.Context
}

type Handler func(ctx ApiContext) error

// 成功的返回
func (c *ApiContext) Success(payload interface{}) error {
	_, err := c.JSON(iris.Map{
		"c": 0,
		"m": "ok",
		"d": payload,
	}, context.JSON{
		StreamingJSON: true,
		UnescapeHTML:  true,
	})
	return err
}

// 失败的返回
func (c *ApiContext) Fail(error error, code int) error {
	_, err := c.JSON(iris.Map{
		"c": code,
		"m": error.Error(),
		"d": nil,
	}, context.JSON{
		StreamingJSON: true,
		UnescapeHTML:  true,
	})
	return err
}

func ApiHandlerWrap(handler Handler) iris.Handler {
	return func(c iris.Context) {
		ctx := ApiContext{c}
		handler(ctx)
	}
}
