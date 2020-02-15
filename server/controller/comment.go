package controller

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
)

func CreateComment(c *model.DBContext) error {
	var comment model.Comment
	// json.NewDecoder(c.Request().Body).Decode(&comment)
	if err := c.Bind(&comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	data, err := model.CreateComment(c.Db, &comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, data)
}

func FetchComments(c *model.DBContext) error {
	records, err := model.FetchComments(c.Db, nil)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, records)
}
