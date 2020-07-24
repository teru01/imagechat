package database

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

type DBContext struct {
	echo.Context
	Db *gorm.DB
}

func ConnectDB(user, password, host, dbname, logMode string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, dbname)
	db, err := gorm.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	db.LogMode(logMode == "1")
	return db
}
