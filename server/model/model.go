package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"
	"os"
	"log"
	"fmt"

	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBContext struct {
	echo.Context
	Db *gorm.DB
}

type Hoge struct {
	gorm.Model
	// id   int `gorm:"type:int;PRIMARY_KEY;AUTO_INCREMENT"`
	name string `gorm:"type:varchar(255)"`
}

func Save(db *gorm.DB, value string) {
	if err := db.Exec("INSERT INTO hoges (name) VALUES (\"hoge\")").Error; err != nil {
		fmt.Fprintf(os.Stderr, "%v", err.Error())
	}
	// db.Create(&Hoge{name: value})
}

func ConnectDB() *gorm.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
