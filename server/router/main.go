package router

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type MainRouter interface {
	setupRouting() *echo.Echo
	StartWebServer()
}

type mainRouter struct {
	authR AuthRouter
}

func NewMainRouter(authR AuthRouter) MainRouter {
	return &mainRouter{authR}
}

func (mainRouter *mainRouter) setupRouting() *echo.Echo {
	router := echo.New()
	router.Use(middleware.CORS())

	mainRouter.authR.SetAuthRouting(router)
	
	return router
}

func (mainRouter *mainRouter) StartWebServer() {
	fmt.Println("Rest API start with echo")

	e := mainRouter.setupRouting()
	e.Logger.Fatal(e.Start(":80"))
}
