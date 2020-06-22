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

	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", "https://localhost:48080", "http://localhost", "https://localhost"},
		AllowMethods:     []string{"GET, DELETE, OPTIONS, POST, PUT"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,Keep-Alive,X-Requested-With,If-Modified-Since"},
	}))

	e.GET("/", handlerWrapper(controller.IndexGet, db))
	e.GET("/posts", handlerWrapper(controller.FetchPosts, db))
	e.GET("/posts/:id", handlerWrapper(controller.FetchPost, db))
	e.POST("/posts", handlerWrapper(controller.RegisterPost, db))

	e.POST("/users", handlerWrapper(controller.SignUp, db))
	e.PUT("/users/:id", handlerWrapper(controller.UpdateUser, db))
	e.DELETE("/users/:id", handlerWrapper(controller.DeleteUser, db))

	e.POST("/comments", handlerWrapper(controller.CreateComment, db))
	e.GET("/comments", handlerWrapper(controller.FetchComments, db))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}

// インタフェースの変換を行う
func handlerWrapper(f func(c *model.DBContext) error, db *gorm.DB) func(echo.Context) error {
	return func(ec echo.Context) error {
		return f(&model.DBContext{ec, db})
	}
}
