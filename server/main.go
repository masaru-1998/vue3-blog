package main

import (
	"sample-table/db"
	"sample-table/utils/validation"
	"sample-table/utils/logic"
	"sample-table/repositories"
	"sample-table/service"
	"sample-table/controller"
	"sample-table/router"
)

func main() {
	db:= db.Init()
	//repositoryパッケージをインスタンス化
	userRepository := repositories.NewUserRepository(db)
	//logicパッケージをインスタンス化
	userLogic := logic.NewUserLogic()
	//validationパッケージをインスタンス化
	authValidation := validation.NewAuthValidation()
	//serviceパッケージをインスタンス化
	authService := service.NewAuthService(authValidation, userRepository, userLogic)
	//controllerパッケージをインスタンス化
	authController := controller.NewAuthController(authService)
	//routerパッケージをインスタンス化
	authRouter := router.NewAuthRouter(authController)
	mainRouter := router.NewMainRouter(authRouter)

	mainRouter.StartWebServer()
}