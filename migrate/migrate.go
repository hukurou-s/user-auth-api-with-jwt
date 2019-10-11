package main

import (
	"fmt"

	"github.com/hukurou-s/user-auth-api-with-jwt/domain"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "user=LEO dbname=user-auth-db password='' sslmode=disable")
	defer db.Close()

	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("success")
	}

	db.AutoMigrate(&domain.User{})
}
