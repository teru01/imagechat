package main

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/"
	"github.com/teru01/image/server/controller"

	"net/http"
	"os"
	"log"
)

func main() {
	db := model.ConnectDB()

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &DBContext
			return next(cc)
		}
	})
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
		AllowOrigins: []string{"*"},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.GET("/", controller.indexGet)
	e.POST("/hoge", controller.handleHoge)
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}


