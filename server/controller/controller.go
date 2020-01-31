package controller

import (
	"net/http"
	"github.com/labstack/echo"
	"github.com/teru01/server/model"

)

func IndexGet(c model.DBContext) error {
	c.String(http.StatusOK, "hello")
}

func HandleHoge(c model.DBContext) error {
	name := c.FormValue("name")
	model.Save(c.db, "name", name)
	return c.String(http.StatusOK, "name: " + name)
}
