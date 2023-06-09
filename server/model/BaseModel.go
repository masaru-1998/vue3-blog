package model

import (
	"time"
)

type BaseModel struct {
	ID uint             `gorm:"primary_key" json: id`
	CreateAt time.Time  `json: createAt`
	UpdateAt time.Time  `json: updateAt`
	DeleteAt *time.Time `json: deleteAt`
}