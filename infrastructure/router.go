package infrastructure

import (
	"io/ioutil"

	jwt "github.com/dgrijalva/jwt-go"
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

	e.POST("/login", func(c echo.Context) error { return controller.Login(c) })

	pubKeyData, err := ioutil.ReadFile("./rsa/id_rsa.pub.pkcs8")
	if err != nil {
		panic(err)
	}

	pubKey, err := jwt.ParseRSAPublicKeyFromPEM(pubKeyData)
	if err != nil {
		panic(err)
	}

	u := e.Group("/user")
	u.Use(middleware.JWTWithConfig(middleware.JWTConfig{
		SigningKey:    pubKey,
		SigningMethod: "RS256",
	}))
	u.POST("/create", func(c echo.Context) error { return controller.Create(c) })

	Echo = e
}
