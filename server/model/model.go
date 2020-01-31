package model

import (
	"github.com/jinzhu/gorm"
	"github.com/labstack/echo"

	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type DBContext {
	echo.Context
	db: *sql.DB
}

type Hoge struct {
	gorm.Model
	id   int
	name string
}

func Save(db *sql.DB, column string, value string) {
	db.Create(&Hoge{name: value})
}

func ConnectDB() *sql.DB {
	dsn := os.Getenv("DSN")
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}
	return db
}
