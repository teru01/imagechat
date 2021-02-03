package controller

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/database"
	"github.com/teru01/image/server/model"
)

func Follow(c *database.DBContext) error {
	var follow model.Follow
	follow.UserID = model.GetAuthSessionData(c, "user_id").(uint)
	i, err := strconv.ParseUint(c.Param("user_id"), 10, 64)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	follow.FollowUserID = uint(i)
	_, err = follow.Create(c.Db)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError)
	}
	return c.NoContent(http.StatusCreated)
}
