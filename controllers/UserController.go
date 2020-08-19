package controllers

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris/v12"
	"maplinking/model"
	"maplinking/services"
)

func Login(ctx iris.Context) {
	name := ctx.PostValue("account")
	password := ctx.PostValue("password")
	user, res := services.Login(name, password)
	fmt.Print(user, res)
	if res {
		//登录成功
		token, _ := services.GenerateJwt(user.Id)
		ctx.JSON(Ret{HTTP_OK, "登录成功", token})
	} else {
		//登录失败
		ctx.JSON(Ret{HTTP_CLIENT_ERROR, "用户名/密码错误，请检查后重试", nil})
	}

}

func GetUserInfo(ctx iris.Context) {
	fmt.Println(ctx.Values())
	userMsg := ctx.Values().Get("jwt").(*jwt.Token).Claims.(jwt.MapClaims)
	id := (int)(userMsg["userid"].(float64))
	if id == 0 {
		ctx.JSON(Ret{
			HTTP_CLIENT_ERROR, "params wrong", nil})
		return
	}
	var user model.User
	if user.GetUserInfo(id) {
		ctx.JSON(Ret{
			HTTP_CLIENT_ERROR, "data not found", nil})
		return
	}
	fmt.Print(user)
	ctx.JSON(Ret{HTTP_OK, "success", user})
	return
}
