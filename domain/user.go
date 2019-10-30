package domain

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Snum     string `gorm:"size:32;unique;not null" json:"snum"`
	Name     string `gorm:"size:255;not null" json:"name"`
	Password string `gorm:"size:255;not null" json:"password"`
}

type LoginParams struct {
	Snum     string `json:"snum"`
	Password string `json:"password"`
}
