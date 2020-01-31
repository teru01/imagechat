package main

import (
	"github.com/labstack/echo"
	// "github.com/labstack/echo/middleware"
	// "database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "github.com/jinzhu/gorm"
	// _ "github.com/jinzhu/gorm/dialects/"
	"github.com/teru01/image/controller"
	"github.com/teru01/image/model"

	// "net/http"
	// "os"
	// "log"
)

func main() {
	db := model.ConnectDB()
	defer db.Close()
	e := echo.New()
	// e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
	// 	return func(c echo.Context) error {
	// 		cc := &model.DBContext{c, db}
	// 		return next(cc)
	// 	}
	// })
	// e.Use(middleware.CORSWithConfig(middleware.CORSConfig {
	// 	AllowOrigins: []string{"*"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPost},
	// }))

	e.GET("/", func(c echo.Context) error {
		return controller.IndexGet(&model.DBContext{c, db})
	})
	e.POST("/hoge", func(c echo.Context) error {
		return controller.HandleHoge(&model.DBContext{c, db})
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}

