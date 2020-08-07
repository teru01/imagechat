package main

import (
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru01/image/server/controller"
	"github.com/teru01/image/server/database"
	customMiddleware "github.com/teru01/image/server/middleware"
	"github.com/teru01/image/server/model"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	db := database.ConnectDB(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"), os.Getenv("QUERY_LOG_MODE"))
	InitializeDB(db)
	defer db.Close()
	e := echo.New()

    api := e.Group("/api")
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: os.Stdout,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", "https://localhost:48080", "http://localhost", "https://localhost", "https://imagechat.ga"},
		AllowMethods:     []string{"GET, DELETE, OPTIONS, POST, PUT"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,Keep-Alive,X-Requested-With,If-Modified-Since"},
	}))
	api.Use(session.Middleware(Store))

	api.GET("/", handlerWrapper(controller.IndexGet, db))
	api.GET("/posts", handlerWrapper(controller.FetchPosts, db))
	api.GET("/posts/:id", handlerWrapper(controller.FetchPost, db))
	api.POST("/posts", handlerWrapper(controller.SubmitPost, db), customMiddleware.SessionAuthentication)

	api.POST("/users", handlerWrapper(controller.SignUp, db))
	// api.PUT("/users/:id", handlerWrapper(controller.UpdateUser, db))
	// api.DELETE("/users/:id", handlerWrapper(controller.DeleteUser, db))

	api.POST("/session", handlerWrapper(controller.Login, db))
	api.GET("/session", handlerWrapper(controller.GetInfo, db), customMiddleware.SessionAuthentication)
	api.DELETE("/session", handlerWrapper(controller.Logout, db), customMiddleware.SessionAuthentication)

	api.POST("/comments", handlerWrapper(controller.CreateComment, db))
	api.GET("/comments", handlerWrapper(controller.FetchComments, db))
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
