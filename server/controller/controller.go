package controller

import (
	"net/http"

	"github.com/teru01/image/server/database"
)

type Err struct {
	Description string `json: "description"`
}

func IndexGet(c *database.DBContext) error {
	return c.String(http.StatusOK, "hello")
}
