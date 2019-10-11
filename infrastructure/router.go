package infrastructure

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"

	"github.com/hukurou-s/user-auth-api-with-jwt/interfaces/controllers"
)

var Echo *echo.Echo

func init() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	controller := controllers.NewUserController(NewSqlHandler())

	e.POST("/user/login", func(c echo.Context) error { return controller.Login(c) })
	e.POST("/user/create", func(c echo.Context) error { return controller.Create(c) })

	Echo = e
}
