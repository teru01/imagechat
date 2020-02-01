package controller

import (
	"net/http"

	"github.com/teru01/image/server/model"
)

func HandleHoge(c *model.DBContext) error {
	h := new(model.HogeForm)
	if err := c.Bind(h); err != nil {
		return c.JSON(http.StatusInternalServerError, &Err{err.Error()})
	}
	if err := model.Insert(c.Db, h.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, &Err{err.Error()})
	}
	return c.NoContent(http.StatusCreated)
}
