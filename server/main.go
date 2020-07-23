package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru01/image/server/controller"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	customMiddleware "github.com/teru01/image/server/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	db := database.ConnectDB()
	InitializeDB(db)
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
	e.Use(session.Middleware(Store))

	e.GET("/", handlerWrapper(controller.IndexGet, db))
	e.GET("/posts", handlerWrapper(controller.FetchPosts, db))
	e.GET("/posts/:id", handlerWrapper(controller.FetchPost, db))
	e.POST("/posts", handlerWrapper(controller.SubmitPost, db), customMiddleware.SessionAuthentication)

	e.POST("/users", handlerWrapper(controller.SignUp, db))
	// e.PUT("/users/:id", handlerWrapper(controller.UpdateUser, db))
	// e.DELETE("/users/:id", handlerWrapper(controller.DeleteUser, db))

	e.POST("/session", handlerWrapper(controller.Login, db))
	e.GET("/session", handlerWrapper(controller.GetInfo, db), customMiddleware.SessionAuthentication)
	e.DELETE("/session", handlerWrapper(controller.Logout, db), customMiddleware.SessionAuthentication)

	e.POST("/comments", handlerWrapper(controller.CreateComment, db))
	e.GET("/comments", handlerWrapper(controller.FetchComments, db))
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8888"))
}

// インタフェースの変換を行う
func handlerWrapper(f func(c *database.DBContext) error, db *gorm.DB) func(echo.Context) error {
	return func(ec echo.Context) error {
		return f(&database.DBContext{ec, db})
	}
}

func InitializeDB(db *gorm.DB) {
	db.AutoMigrate(&model.Post{})
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Comment{})
}
