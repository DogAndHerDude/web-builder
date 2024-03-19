package user

import (
	"net/http"

	custom_middleware "app/middleware"
	"app/utils"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	userService IUserService
}

type GetMeResult struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *UserHandlers) GetMeHandler(c echo.Context) error {
	claims, ok := c.Get("user").(utils.Claims)
	if !ok {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	user, err := h.userService.GetUserByID(claims.ID)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound, "Not found")
	}

	c.JSON(http.StatusFound, struct {
		Data GetMeResult `json:"data"`
	}{
		Data: GetMeResult{
			user.ID,
			user.Email,
		},
	})

	return nil
}

func RegisterHandlers(s IUserService, e *echo.Group) {
	group := e.Group("/user")
	handlers := UserHandlers{
		userService: s,
	}

	group.Use(custom_middleware.NewAuthorizeMiddleware())
	group.GET("/", handlers.GetMeHandler)
}
