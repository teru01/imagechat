package controller

import (
	"net/http"

	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
	// "github.com/gorilla/sessions"
)

func Login(c *database.DBContext) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	_, err := model.NewSession(&user, c)
	if err != nil {
		return echo.NewHTTPError(http.StatusUnauthorized, err.Error())
	}
	// return c.JSON(http.StatusOK, User{})
	return c.NoContent(http.StatusOK)
}

func GetInfo(c *database.DBContext) error {
	sess, err := session.Get("auth", c)
	if err != nil {
		return err
	}
	ret := struct {
		Name string
	}{
		Name: sess.Values["name"].(string),
	}
	return c.JSON(http.StatusOK, ret)
}
