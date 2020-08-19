package services

import (
	"github.com/dgrijalva/jwt-go"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"github.com/kataras/iris/v12/context"
	"maplinking/model"
	"time"
)

const SecretKey = "this is my maplinking"

//生成jwt
func GenerateJwt(userid int) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := make(jwt.MapClaims)
	claims["userid"] = userid
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(24)).Unix()
	claims["iat"] = time.Now().Unix()
	token.Claims = claims
	tokenString, err := token.SignedString([]byte(SecretKey))
	return tokenString, err
}

func test() {

}

func ValidateJwt() *jwtmiddleware.Middleware {
	jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
		//这个方法将验证jwt的token
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(SecretKey), nil
		},
		//加密的方式
		SigningMethod: jwt.SigningMethodHS256,
		ErrorHandler: func(ctx context.Context, err error) {
			if err == nil {
				return
			} else {
				ctx.StatusCode(200)
				ctx.JSON(map[string]string{"code": "401", "message": "身份验证失败,请重新登录", "data": ""})
			}
		},
	})
	return jwtHandler
}

//验证登录
func Login(name string, password string) (model.User, bool) {
	var user model.User
	if user.GetUserByLogin(name, password) {
		return user, false
	}
	return user, true
}
