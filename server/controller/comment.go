package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
)

func CreateComment(c *database.DBContext) error {
	var comment model.Comment
	if err := c.Bind(&comment); err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	createdComment, err := comment.Create(c.Db, &comment)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusCreated, createdComment)
}

func FetchComments(c *database.DBContext) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}
	comment := model.Comment{}
	records, err := comment.Select(c.Db, &map[string]interface{}{"post_id": c.Param("post_id")}, offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, records)
}

// func FetchComment(c *database.DBContext) error {
// 	records, err := model.FetchComments(c.Db, nil)
// 	if err != nil {
// 		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
// 	}
// 	return c.JSON(http.StatusOK, records)
// }
