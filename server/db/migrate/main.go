package main

import (
	"sample-table/db"
	"sample-table/model"

	"gorm.io/gorm"
)

func migrate(dbCon *gorm.DB) {
	dbCon.AutoMigrate(&model.User{})
}

func main() {
	dbCon := db.Init()
	defer db.CloseDB(dbCon)

	//migrage実行
	migrate(dbCon)
}