package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/form"
	"github.com/teru01/image/server/model"
)

func SignUp(c *database.DBContext) error {
	var userForm form.UserForm
	if err := c.Bind(&userForm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	user := model.User{}
	if err := user.CreateUser(c.Db, userForm); err != nil {
		echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	return c.NoContent(http.StatusCreated)
}

// func UpdateUser(c *database.DBContext) error {
// 	var user model.User
// 	_id, err := strconv.Atoi(c.Param("id"))
// 	id := uint(_id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	user.ID = id
// 	var m map[string]interface{} = map[string]interface{}{}
// 	if err := c.Bind(&m); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	if _, ok := m["password"]; ok {
// 		// password :=
// 		hashed, err := hashPassword(m["password"].(string))
// 		if err != nil {
// 			return err
// 		}
// 		m["password"] = hashed
// 	}
// 	data, err := model.UpdateUser(c.Db, &user, m)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, data)
// }

// func DeleteUser(c *database.DBContext) error {
// 	var user model.User
// 	_id, err := strconv.Atoi(c.Param("id"))
// 	id := uint(_id)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
// 	}
// 	user.ID = id
// 	if err = model.DeleteUser(c.Db, &user); err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, user)
// }
