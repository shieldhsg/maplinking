//路由注册组件
package controllers

import (
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12"
	ws "maplinking/websocket"
)

//返回码
const HTTP_OK = 200
const HTTP_CLIENT_ERROR = 201
const SERVER_ERROR = 500

//返回结构体
type Ret struct {
	Code    int
	Message string
	Data    interface{}
}

func Register(app *iris.Application, j *jwtmiddleware.Middleware) {
	//给所有的方法一个option处理 否则可能会出现跨域
	common := app.Party("/")
	{
		common.Options("*", func(ctx iris.Context) {
			ctx.Next()
		})
	}
	userController := app.Party("/user") //用户相关handler
	app.Get("/ws", ws.ServeWs)           //websocket相关handler
	//不需要登录的接口
	userController.Post("/login", Login)
	//需要登录的接口
	//websocket相关接口
	userController.Get("/info", j.Serve, GetUserInfo)
}
