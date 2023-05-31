package controller

import (
	"net/http"

	"sample-table/service"


	"github.com/labstack/echo/v4"
)

//インターフェースの設定
type AuthController interface {
	SignUp(c echo.Context) error
}
//コントローラ構造体を作成
type authController struct {
	as service.AuthService
}
//構造体をインスタンス化
func NewAuthController(as service.AuthService) AuthController{
	return &authController{as}
}

//サインアップする処理
func (ac *authController) SignUp(c echo.Context) error {
	//サインアップ処理を実行
	createUser, err := ac.as.SignUp(c)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	} 

	responseBody, err := ac.as.CreateAuthResponse(&createUser, http.StatusCreated)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, responseBody)
}