package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/form"
	"github.com/teru01/image/server/model"
)

func FetchPosts(c *database.DBContext) error {
	offset, err := strconv.Atoi(c.QueryParam("offset"))
	if err != nil {
		offset = 0
	}

	limit, err := strconv.Atoi(c.QueryParam("limit"))
	if err != nil {
		limit = 20
	}

	posts, err := model.SelectPosts(c.Db, nil, offset, limit)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.JSON(http.StatusOK, posts)
}

func FetchPost(c *database.DBContext) error {
	posts, err := model.SelectPosts(c.Db, &map[string]interface{}{"id": c.Param("id")}, 0, 1)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}

	if len(posts) == 0 {
		return echo.NewHTTPError(http.StatusNotFound, "Not Found")
	}

	return c.JSON(http.StatusOK, posts[0])
}

func SubmitPost(c *database.DBContext) error {
	fileHeader, err := c.FormFile("photo")
	if err != nil {
		return err
	}

	var postForm form.PostForm
	postForm.Name = c.FormValue("name")

	if err := model.SubmitPost(c.Db, fileHeader, postForm); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
	}
	return c.NoContent(http.StatusCreated)
}
