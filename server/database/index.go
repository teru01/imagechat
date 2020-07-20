package database

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/labstack/echo/v4"
)

type DBContext struct {
	echo.Context
	Db *gorm.DB
}

func ConnectDB() *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", "root", os.Getenv("MYSQL_ROOT_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"))
	db, err := gorm.Open("mysql", dsn+"?parseTime=true")
	if err != nil {
		log.Fatal(err)
	}
	return db
}


