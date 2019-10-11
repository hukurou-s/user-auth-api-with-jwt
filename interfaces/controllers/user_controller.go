package controllers

import (
	_ "encoding/hex"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"

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
		return c.JSON(http.StatusInternalServerError, err)
	}

	user, err := controller.Interactor.UserBySnum(loginParams.Snum)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	if compareHashedPassword(user.Password, loginParams.Password) {
		return c.JSON(http.StatusOK, struct {
			Status string `json:"status"`
		}{
			Status: "success",
		})
	}
	return c.JSON(http.StatusUnauthorized, struct {
		Status string `json:"status"`
	}{
		Status: "fail",
	})
}

func toHashPassword(pass string) string {
	converted, _ := bcrypt.GenerateFromPassword([]byte(pass), bcrypt.DefaultCost)
	return string(converted) //hex.EncodeToString(converted[:])
}

func compareHashedPassword(hash string, pass string) bool {
	fmt.Println(hash + "\n")
	fmt.Println(pass + "\n")
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(pass))
	if err == nil {
		return true
	}
	return false
}
