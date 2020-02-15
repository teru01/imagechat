package controller

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
	"golang.org/x/crypto/bcrypt"
)


func SignUp(c *model.DBContext) error {
	var user model.User
	if err := c.Bind(&user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	hashed, err := hashPassword(user.Password)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	user.Password = hashed
	form, err := model.CreateUser(c.Db, &user)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, form)
}

func hashPassword(original string) (string, error) {
	hashedPasswd, err := bcrypt.GenerateFromPassword([]byte(original), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hashedPasswd), nil
}

func UpdateUser(c *model.DBContext) error {
	var user model.User
	_id, err := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.ID = id
	var m map[string]interface{} = map[string]interface{}{}
	if err := c.Bind(&m); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if _, ok := m["password"]; ok {
		// password := 
		hashed, err := hashPassword(m["password"].(string))
		if err != nil {
			return err
		}
		m["password"] = hashed
	}
	data, err := model.UpdateUser(c.Db, &user, m)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, data)
}

func DeleteUser(c *model.DBContext) error {
	var user model.User
	_id, err := strconv.Atoi(c.Param("id"))
	id := uint(_id)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	user.ID = id
	if err = model.DeleteUser(c.Db, &user); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, user)
}
