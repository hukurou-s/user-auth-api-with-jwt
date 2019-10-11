package infrastructure

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"

	"github.com/hukurou-s/user-auth-api-with-jwt/interfaces/database"
)

func NewSqlHandler() database.SqlHandler {
	db, err := gorm.Open("postgres", "user=LEO dbname=user-auth-db password='' sslmode=disable")

	if err != nil {
		panic(err)
	}
	return db
}
