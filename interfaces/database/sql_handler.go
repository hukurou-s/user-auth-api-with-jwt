package database

import "github.com/jinzhu/gorm"

type SqlHandler interface {
	Create(interface{}) *gorm.DB
	Save(interface{}) *gorm.DB
	Where(interface{}, ...interface{}) *gorm.DB
	First(interface{}, ...interface{}) *gorm.DB
}
