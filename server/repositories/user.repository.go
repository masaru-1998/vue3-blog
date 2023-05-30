package repositories

import (
	"gorm.io/gorm"
	"fmt"

	"sample-table/model"
)

type UserRepository interface {
	CreateUser(createUser *model.User) error
	GetUserByEmail(user *model.User, email string) error
	GetUserAllByEmail(user *[]model.User, email string) error 
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db}
}

/*
ユーザ取得
*/
func (ur *userRepository) GetUserByEmail(user *model.User, email string) error {
	if err := ur.db.Where("email=?", email).Find(&user).Error; err != nil {
		return fmt.Errorf("ユーザ取得エラー: %v", err)
	}

	return nil
}

/*
ユーザー一覧取得
*/
func (ur *userRepository) GetUserAllByEmail(user *[]model.User, email string) error {
	if err := ur.db.Where("email=?", email).Find(&user).Error; err != nil {
		return fmt.Errorf("ユーザ取得エラー: %v", err)
	}

	return nil
}

/*
ユーザの登録
*/
func (ur *userRepository) CreateUser(createUser *model.User) error {
	if err := ur.db.Create(&createUser).Error; err != nil {
		return fmt.Errorf("ユーザ作成失敗: %v", err)
	}

	return nil
}
