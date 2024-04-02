package custom_middleware

import (
	"net/http"

	"github.com/DogAndHerDude/web-builder/utils"

	"github.com/labstack/echo/v4"
)

func NewAuthorizeMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			if cookie, err := c.Cookie("jwt"); err != nil {
				c.String(http.StatusUnauthorized, "Unauthorized")
			} else {
				token := cookie.String()
				if claims, err := utils.VerifyJWT(token); err != nil {
					c.String(http.StatusUnauthorized, "Unauthorized")
				} else {
					c.Set("user", claims)

					return next(c)
				}
			}

			return nil
		}
	}
}
