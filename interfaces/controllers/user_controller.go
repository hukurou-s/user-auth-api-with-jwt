package controllers

import (
	_ "fmt"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"net/http"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hukurou-s/user-auth-api-with-jwt/domain"
	"github.com/hukurou-s/user-auth-api-with-jwt/interfaces/database"
	"github.com/hukurou-s/user-auth-api-with-jwt/usecase"
)

type UserController struct {
	Interactor usecase.UserInteractor
}

func NewUserController(sqlHandler database.SqlHandler) *UserController {
	return &UserController{
		Interactor: usecase.UserInteractor{
			UserRepository: &database.UserRepository{
				SqlHandler: sqlHandler,
			},
		},
	}
}

func (controller *UserController) Create(c Context) error {
	user := domain.User{}
	if err := c.Bind(&user); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	if user.Name == "" || user.Snum == "" || user.Password == "" {
		return c.JSON(http.StatusBadRequest, struct {
			Status string `json:"status"`
		}{
			Status: "fail",
		})
	}

	user.Password = toHashPassword(user.Password)

	err := controller.Interactor.Add(user)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
	}{
		Status: "success",
	})
}

func (controller *UserController) Login(c Context) error {
	loginParams := new(domain.LoginParams)
	if err := c.Bind(loginParams); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	user, err := controller.Interactor.UserBySnum(loginParams.Snum)
	if err != nil {
		return c.JSON(http.StatusUnauthorized, err)
	}

	if !compareHashedPassword(user.Password, loginParams.Password) {
		return c.JSON(http.StatusUnauthorized, struct {
			Status string `json:"status"`
		}{
			Status: "fail",
		})

	}

	keyData, err := ioutil.ReadFile("./rsa/id_rsa")
	if err != nil {
		panic(err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyData)
	if err != nil {
		panic(err)
	}

	token := jwt.New(jwt.SigningMethodRS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["sub"] = user.ID
	claims["snum"] = user.Snum
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	t, err := token.SignedString(key)
	if err != nil {
		panic(err)
	}

	return c.JSON(http.StatusOK, struct {
		Status string `json:"status"`
		Token  string `json:"token"`
	}{
		Status: "success",
		Token:  t,
	})
}

func toHashPassword(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(converted)
}

func compareHashedPassword(hash string, pass string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
