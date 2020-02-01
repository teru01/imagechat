package model

import (
	"log"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type DBContext struct {
	echo.Context
	Db *gorm.DB
}

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	db.AutoMigrate(&Hoge{})
	return db
}
