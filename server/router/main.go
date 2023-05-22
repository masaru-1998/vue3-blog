package router

import (
	"fmt"
	"net/http"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)



func StartWebServer() {
	e := echo.New()

	e.Use(middleware.CORS())
	e.POST("/signup", func(c echo.Context) error {
		firstName := c.FormValue("firstName")
		lastName := c.FormValue("lastName")
		email := c.FormValue("email")
		password := c.FormValue("password")
		
		response := fmt.Sprintf("firstName: %s, lastName: %s, email: %s, password: %s", firstName, lastName, email, password)
		return c.String(http.StatusOK, response)
	})
	e.Logger.Fatal(e.Start(":80"))
}
