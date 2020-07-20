package controller

import (
	"net/http"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
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
	return c.NoContent(http.StatusOK)
}

func Logout(c *database.DBContext) error {
	sess, err := session.Get(model.SessionName, c)
	if err != nil {
		return err
	}
	sess.Options.MaxAge = -1
	sess.Save(c.Request(), c.Response())
	return nil
}

func GetInfo(c *database.DBContext) error {
	sess, err := session.Get(model.SessionName, c)
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
