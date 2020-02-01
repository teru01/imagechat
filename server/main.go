package main

import (
	"fmt"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru01/image/server/controller"
	"github.com/teru01/image/server/model"
	"github.com/jinzhu/gorm"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db := model.ConnectDB()
	defer db.Close()
	e := echo.New()

	e.Use(middleware.BodyDump(func(c echo.Context, reqBody, resBody []byte) {
		fmt.Fprintf(os.Stderr, "request: %v\n", string(reqBody))
	}))
	e.GET("/", handlerWrapper(controller.IndexGet, db))

	e.POST("/hoge", handlerWrapper(controller.HandleHoge, db))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}


// インタフェースの変換を行う
func handlerWrapper(f func (c *model.DBContext) error, db *gorm.DB) (func (echo.Context) error) {
	return func(ec echo.Context) error {
		return f(&model.DBContext{ec, db})
	}
}
