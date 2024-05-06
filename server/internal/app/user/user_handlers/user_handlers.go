package user_handlers

import (
	"net/http"

<<<<<<< HEAD:server/internal/app/user/user_handlers/user_handlers.go
	"github.com/DogAndHerDude/web-builder/internal/app/user/user_service"
=======
>>>>>>> 973db0e (chore(project): udpate imports and add build targets):server/user/handlers.go
	"github.com/DogAndHerDude/web-builder/internal/pkg/jwt_utils"
	custom_middleware "github.com/DogAndHerDude/web-builder/middleware"

	"github.com/labstack/echo/v4"
)

type UserHandlers struct {
	userService user_service.IUserService
}

type GetMeResult struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

func (h *UserHandlers) GetMeHandler(c echo.Context) error {
	claims, ok := c.Get("user").(jwt_utils.Claims)
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

func RegisterHandlers(s user_service.IUserService, e *echo.Group) {
	group := e.Group("/user")
	handlers := UserHandlers{
		userService: s,
	}

	group.Use(custom_middleware.NewAuthorizeMiddleware())
	group.GET("/", handlers.GetMeHandler)
}
