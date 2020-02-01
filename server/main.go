package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru01/image/server/controller"
	"github.com/teru01/image/server/model"

	_ "github.com/go-sql-driver/mysql"
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

