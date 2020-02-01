package controller

import (
	"net/http"

	"github.com/teru01/image/server/model"
)

type Err struct {
	Description string `json: "description"`
}

func IndexGet(c *model.DBContext) error {
	return c.String(http.StatusOK, "hello")
}
