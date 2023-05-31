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
	CreateAuthResponse(user *model.User, code int) (model.AuthResponse, error)
}

type authService struct {
	av validation.AuthValidation
	ur repositories.UserRepository
	ul logic.UserLogic
	jl logic.JWTTokenLogic
}

func NewAuthService(av validation.AuthValidation, ur repositories.UserRepository, ul logic.UserLogic, jl logic.JWTTokenLogic) AuthService {
	return &authService{av, ur, ul, jl}
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

func (as *authService) CreateAuthResponse(user *model.User, code int) (model.AuthResponse, error) {
	//jwtトークンを作成
	jwtToken, err := as.jl.CreateJWTToken(user)
	if err != nil {
		return model.AuthResponse{}, err
	}
	//レスポンスデータを生成
	var responseBody model.AuthResponse
	responseBody.Token = jwtToken
	responseBody.User.Name = user.Name
	responseBody.User.Email = user.Email
	responseBody.User.BaseModel.ID = user.ID
	responseBody.User.BaseModel.CreateAt = user.CreateAt
	responseBody.User.BaseModel.UpdateAt = user.UpdateAt
	responseBody.User.BaseModel.DeleteAt = user.DeleteAt
	
	return responseBody, nil
}