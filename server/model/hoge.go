package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"

	_ "github.com/jinzhu/gorm/dialects/mysql"
)


type HogeForm struct {
	Name string `json: "name" gorm:"type:varchar(255)"`
}

type Hoge struct {
	gorm.Model
	HogeForm
}

func Insert(db *gorm.DB, value string) error {
	return db.Create(&Hoge{HogeForm: HogeForm{value}}).Error
}
