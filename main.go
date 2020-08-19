package main

import (
	"flag"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"maplinking/controllers"
	"maplinking/services"
	ws "maplinking/websocket"
)

//跨域
func Cors(ctx iris.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	if ctx.Request().Method == "OPTIONS" {
		ctx.Header("Access-Control-Allow-Methods", "GET,POST,PUT,DELETE,PATCH,OPTIONS")
		ctx.Header("Access-Control-Allow-Headers", "Content-Type, Accept, Authorization")
		ctx.StatusCode(200)
		return
	}
	ctx.Next()
}

func test(ctx iris.Context) {
	//如果解密成功，将会进入这里,获取解密了的token
	token := ctx.Values().Get("jwt").(*jwt.Token)
	//或者这样

	ctx.Writef("This is an authenticated request\n")
	ctx.Writef("Claim content:\n")
	//可以了解一下token的数据结构
	ctx.Writef("%s", token.Signature)
}

func main() {
	app := iris.New()
	app.Use(Cors)
	j := services.ValidateJwt()
	controllers.Register(app, j) //注册路由
	//启动websocket
	flag.Parse()
	ws.HubInstance = ws.NewHub()
	go ws.HubInstance.Run()
	app.Logger().SetLevel("debug")
	app.Run(iris.Addr(":8080"), iris.WithoutServerError(iris.ErrServerClosed))
	a := 1
	fmt.Println(a)
}
