package service

import (
	"sample-table/utils/validation"
	"sample-table/model"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	SignUp(c echo.Context) (model.SignUpRequest, error)
}

type authService struct {
	av validation.AuthValidation
}

func NewAuthService(av validation.AuthValidation) AuthService {
	return &authService{av}
}

func (as *authService) SignUp(c echo.Context) (model.SignUpRequest, error){
	//パラメータを取得
	LastName := c.FormValue("lastName")
	FirstName := c.FormValue("firstName")
	Email := c.FormValue("email")
	Password := c.FormValue("password")
	signUpRequestParams := model.SignUpRequest{
		FirstName: FirstName,
		LastName: LastName,
		Email: Email,
		Password: Password,
	}
	//パラメータのバリデーション
	if err := as.av.SignUpValidate(signUpRequestParams); err != nil {
		return signUpRequestParams, err
	}

	return signUpRequestParams, nil
}

