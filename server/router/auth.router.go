package router

import (
	"sample-table/controller"
	"github.com/labstack/echo/v4"
)

//インターフェースの作成
type AuthRouter interface {
	SetAuthRouting(router *echo.Echo)
}
//auth構造体の作成
type authRouter struct {
	ac controller.AuthController
}
//構造体のインスタンス
func NewAuthRouter(ac controller.AuthController) AuthRouter{
	return &authRouter{ac}
}
//authルーティングを作成
func (ar *authRouter) SetAuthRouting(router *echo.Echo) {
	router.POST("/api/signup", ar.ac.SignUp)
}
