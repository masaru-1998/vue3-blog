package router

import (
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func StartWebServer() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.POST("/signup", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.Logger.Fatal(e.Start(":80"))
}
