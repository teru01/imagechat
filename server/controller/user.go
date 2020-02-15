package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
	"golang.org/x/crypto/bcrypt"
)


func SignUp(c *model.DBContext) error {
	var userForm model.UserForm
	if err := c.Bind(&userForm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(userForm.Password), bcrypt.DefaultCost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	userForm.Password = fmt.Sprintf("%x", hashedPasswd)
	form, err := model.Create(c.Db, &userForm)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, form)
}
