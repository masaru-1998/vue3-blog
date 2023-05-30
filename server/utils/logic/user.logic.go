package logic 

import "golang.org/x/crypto/bcrypt"

type UserLogic interface {
	ChangeHashPassword(password string) []byte
}

type userLogic struct {

}

func NewUserLogic() UserLogic {
	return &userLogic{}
}

func (ul *userLogic) ChangeHashPassword(password string) []byte {
	hassPass, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return hassPass
}