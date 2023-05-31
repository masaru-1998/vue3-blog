package logic

import (
	"fmt"
	"time"
	"os"

	"github.com/dgrijalva/jwt-go"
	"github.com/joho/godotenv"

	"sample-table/model"
)

type JWTTokenLogic interface {
	CreateJWTToken(user *model.User) (string, error)
}

type jwtTokenLogic struct {}

func NewJWTLogic() JWTTokenLogic {
	return &jwtTokenLogic{}
}

type UserClaims struct {
	Name string `json: name`
	Email string `json: email`
	jwt.StandardClaims
}

func (jl *jwtTokenLogic) CreateJWTToken(user *model.User) (string, error) {
	//クレームの設定
	//クレームはtokenの一部で、トークンに付加情報を与えることができるもののこと
	claims := &UserClaims{
		Name:  user.Name,
		Email: user.Email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	//トークンを生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//環境変数読み込み
	if err := godotenv.Load("./.env"); err != nil {
		fmt.Errorf("環境変数の読み込みに失敗しました %v", err)
		return "", err
	}
	//トークンを文字列として作成
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_KEY")))
	if err != nil {
		fmt.Errorf("トークンを文字列変換するのに失敗しました %v", err)
		return "" , err
	}

	return tokenString, nil
}
