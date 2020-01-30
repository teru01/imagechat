package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "hoge, world")
	})
	e.Static("/static", "static")
	e.Logger.Fatal(e.Start(":8080"))
}
