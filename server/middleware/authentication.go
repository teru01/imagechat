package middleware

import 	(
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo-contrib/session"

)

func SessionAuthentication(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		session, _ := session.Get("auth", c)
		if session.IsNew {
			return echo.ErrUnauthorized
		}
		return next(c)
	}
}
