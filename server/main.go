package main

import (
	"os"

	"github.com/jinzhu/gorm"
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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig {
		Output: os.Stdout,
	}))

	e.GET("/", handlerWrapper(controller.IndexGet, db))
	e.GET("/hoges", handlerWrapper(controller.FetchHoges, db))
	e.GET("/hoges/:id", handlerWrapper(controller.FetchHoge, db))
	e.POST("/hoges", handlerWrapper(controller.RegisterHoge, db))
	e.POST("/users", handlerWrapper(controller.SignUp, db))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}


// インタフェースの変換を行う
func handlerWrapper(f func (c *model.DBContext) error, db *gorm.DB) (func (echo.Context) error) {
	return func(ec echo.Context) error {
		return f(&model.DBContext{ec, db})
	}
}

