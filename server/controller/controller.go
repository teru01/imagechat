package controller

import (
	"net/http"

	"github.com/teru01/image/server/model"
)

func IndexGet(c *model.DBContext) error {
	return c.String(http.StatusOK, "hello")
}

func HandleHoge(c *model.DBContext) error {
	name := c.FormValue("name")
	model.Save(c.Db, name)
	return c.String(http.StatusOK, "name: " + name)
}
