package user

import "github.com/labstack/echo/v4"

type UserHandlers struct {
	s IUserService
}

func (h *UserHandlers) GetMeHandler(c echo.Context) error {
	return nil
}

func RegisterHandlers(s IUserService, e *echo.Group) {
	handlers := UserHandlers{
		s: s,
	}

	e.GET("/user", handlers.GetMeHandler)
}
