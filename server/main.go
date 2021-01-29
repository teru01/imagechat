package main

import (
	"io"
	"net"
	"net/http"
	"os"

	"github.com/jinzhu/gorm"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/teru01/image/server/controller"
	"github.com/teru01/image/server/database"
	customMiddleware "github.com/teru01/image/server/middleware"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
)

var Store = sessions.NewCookieStore([]byte(os.Getenv("SESSION_KEY")))

func main() {
	db := database.ConnectDB(os.Getenv("MYSQL_USER"), os.Getenv("MYSQL_PASSWORD"), os.Getenv("MYSQL_HOST"), os.Getenv("MYSQL_DATABASE"), os.Getenv("QUERY_LOG_MODE"))
	defer db.Close()
	e := echo.New()

	api := e.Group("/api")

	var output io.Writer = os.Stdout
	raddr, err := net.ResolveUDPAddr("udp", "vector-server:50000")
	if err == nil {
		conn, err := net.DialUDP("udp", nil, raddr)
		if err == nil {
			output = io.MultiWriter(output, conn)
		}
	}
	api.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Output: output,
	}))
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:8080", "https://localhost:48080", "http://localhost", "https://localhost", "https://imagechat.ga"},
		AllowMethods:     []string{"GET, DELETE, OPTIONS, POST, PUT"},
		AllowCredentials: true,
		AllowHeaders:     []string{"Authorization,Content-Type,Accept,Origin,User-Agent,DNT,Cache-Control,X-Mx-ReqToken,Keep-Alive,X-Requested-With,If-Modified-Since"},
	}))
	api.Use(session.Middleware(Store))
	api.GET("/health", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})
	api.GET("/", handlerWrapper(controller.IndexGet, db))
	api.GET("/posts", handlerWrapper(controller.FetchPosts, db))
	api.GET("/posts/:id", handlerWrapper(controller.FetchPost, db))
	api.POST("/posts", handlerWrapper(controller.SubmitPost, db), customMiddleware.SessionAuthentication)

	api.POST("/users", handlerWrapper(controller.SignUp, db))
	api.GET("/users/:id", handlerWrapper(controller.FetchUser, db))
	// api.PUT("/users/:id", handlerWrapper(controller.UpdateUser, db))
	// api.DELETE("/users/:id", handlerWrapper(controller.DeleteUser, db))

	api.POST("/session", handlerWrapper(controller.Login, db))
	api.GET("/session", handlerWrapper(controller.GetInfo, db), customMiddleware.SessionAuthentication)
	api.DELETE("/session", handlerWrapper(controller.Logout, db), customMiddleware.SessionAuthentication)

	api.POST("/:post_id/comments", handlerWrapper(controller.CreateComment, db))
	api.GET("/:post_id/comments", handlerWrapper(controller.FetchComments, db))
	// api.GET("/:post_id/comments/:id", handlerWrapper(controller.FetchComment, db))
	e.Logger.Fatal(e.Start(":8888"))
}

// インタフェースの変換を行う
func handlerWrapper(f func(c *database.DBContext) error, db *gorm.DB) func(echo.Context) error {
	return func(ec echo.Context) error {
		return f(&database.DBContext{ec, db})
	}
}
