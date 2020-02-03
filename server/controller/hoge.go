package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
)

func RegisterHoge(c *model.DBContext) error {
	h := new(model.HogeForm)
	if err := c.Bind(h); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	if err := model.Insert(c.Db, h.Name); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}

func FetchHoges(c *model.DBContext) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	hoges, err := model.HogeSelect(c.Db, nil, offset, limit)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, hoges)
}

func FetchHoge(c *model.DBContext) error {
	hoges, err := model.HogeSelect(c.Db, &map[string]interface{}{"id": c.Param("id")}, 0, 1)
	if err != nil{
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(hoges) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, hoges[0])
}

