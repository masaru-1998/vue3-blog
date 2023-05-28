package main

import (
	"sample-table/utils/validation"
	"sample-table/service"
	"sample-table/controller"
	"sample-table/router"
)

func main() {
	authValidation := validation.NewAuthValidation()

	authService := service.NewAuthService(authValidation)

	authController := controller.NewAuthController(authService)

	authRouter := router.NewAuthRouter(authController)
	mainRouter := router.NewMainRouter(authRouter)

	mainRouter.StartWebServer()
}