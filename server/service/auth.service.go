package service

import (
	"errors"

	"sample-table/utils/validation"
	"sample-table/utils/logic"
	"sample-table/model"
	"sample-table/repositories"
	"github.com/labstack/echo/v4"
)

type AuthService interface {
	SignUp(c echo.Context) (model.User, error)
}

type authService struct {
	av validation.AuthValidation
	ur repositories.UserRepository
	ul logic.UserLogic
}

func NewAuthService(av validation.AuthValidation, ur repositories.UserRepository, ul logic.UserLogic) AuthService {
	return &authService{av, ur, ul}
}

func (as *authService) SignUp(c echo.Context) (model.User, error){
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
		return model.User{}, err
	}
	//既存アカウントが存在していないかの確認
	var users []model.User
	if err := as.ur.GetUserAllByEmail(&users, Email); err != nil {
		return model.User{}, err
	}
	if(len(users) != 0) {
		return model.User{}, errors.New("すでにアカウントが存在します")
	}
	//アカウント情報をデータベースに保存
	var createUser model.User
	hassPass := as.ul.ChangeHashPassword(Password)
	createUser.Name = LastName + " " + FirstName
	createUser.Email = Email
	createUser.Password = string(hassPass)

	if err := as.ur.CreateUser(&createUser); err != nil {
		return model.User{}, err
	}
	

	return createUser, nil
}

