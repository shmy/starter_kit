package main

import (
	"fmt"
	"github.com/kataras/iris"
	"github.com/kataras/iris/core/errors"
	"github.com/lexkong/log"
	"github.com/spf13/viper"
	"net/http"
	"starter_kit/config"
	"starter_kit/model"
	"starter_kit/routing"
	"starter_kit/util"
	"time"
)

func init() {
	// 初始化配置工具和日志
	config.Init("")
	// 初始化数据库
	model.Init()
}
func main() {
	app := iris.Default()
	app.OnAnyErrorCode(func(c iris.Context) {
		ctx := util.ApiContext{c}
		ctx.Fail(errors.New("server error"), ctx.GetStatusCode())
	})
	routing.RouteMapping(app)
	host := viper.GetString("server.host")
	port := viper.GetString("server.port")
	url := fmt.Sprintf("%s:%s", host, port)
	// 启动时自检
	go func() {
		if err := pingServer(url); err != nil {
			log.Fatal("The router has no response, or it might took too long to start up.", err)
		}
		log.Info("The router has been deployed successfully.")
	}()

	app.Run(iris.Addr(url))

}

// 自己Ping自己
func pingServer(url string) error {
	url += "/svcd/health"
	for i := 0; i < 20; i++ {
		resp, err := http.Get("http://" + url)
		if err == nil && resp.StatusCode == http.StatusOK {
			return nil
		}
		log.Info("Waiting for the router, retry in 1 second.")
		time.Sleep(time.Second)
	}
	return errors.New("Cannot connect to the router.")
}
