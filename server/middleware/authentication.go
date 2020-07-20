package middleware

import (
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/teru01/image/server/model"
)

func SessionAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := session.Get(model.SessionName, c)
		if session.IsNew {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
