package auth_handlers

import (
	"net/http"

	"github.com/DogAndHerDude/web-builder/internal/app/auth/auth_service"
	"github.com/DogAndHerDude/web-builder/internal/app/user/user_service"
	hash_utils "github.com/DogAndHerDude/web-builder/internal/pkg/hash_utils"
	"github.com/DogAndHerDude/web-builder/internal/pkg/jwt_utils"

	"github.com/labstack/echo/v4"
)

// Needs data avalidation
// Needs specific password complexity & email validity logic
type SignupPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type LoginPayload struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required,min=8,max=40"`
}

type AuthHandlers struct {
	userService user_service.IUserService
	authService auth_service.IAuthService
}

func (h *AuthHandlers) Signup(c echo.Context) error {
	payload := new(SignupPayload)

	if err := c.Bind(payload); err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}
	if err := c.Validate(payload); err != nil {
		return nil
	}

	salt, err := hash_utils.RandomSecret(10)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	hashedPassword, err := hash_utils.HashString(payload.Password, salt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	user, err := h.userService.CreateUser(payload.Email, string(salt), string(hashedPassword))
	if err != nil {
		return echo.NewHTTPError(http.StatusExpectationFailed)
	}

	token, err := h.authService.GenerateJWT(auth_service.ClaimValues{
		ID: user.ID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	c.SetCookie(&http.Cookie{
		Name:     "jwt",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})
	c.NoContent(http.StatusCreated)

	return nil
}

func (h *AuthHandlers) Authenticate(c echo.Context) error {
	var payload LoginPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotAcceptable, "Invalid payload")
	}

	user, err := h.userService.GetUserByEmail(payload.Email)
	if err != nil {
		return echo.NewHTTPError(http.StatusNotFound)
	}

	c.JSON(http.StatusOK, user)

	return nil
}

func (h *AuthHandlers) Authorize(c echo.Context) error {
	_, ok := c.Get("user").(jwt_utils.Claims)
	if !ok {
		// TODO: remove cookie
		return echo.NewHTTPError(http.StatusUnauthorized)
	}

	c.NoContent(http.StatusOK)

	return nil
}

func RegisterHandlers(e *echo.Group, u user_service.IUserService, a auth_service.IAuthService) {
	h := &AuthHandlers{
		userService: u,
		authService: a,
	}
	subGroup := e.Group("/auth")
	subGroup.POST("/signup", h.Signup)
	subGroup.POST("/authenticate", h.Authenticate)
	subGroup.GET("/authorize", h.Authorize)
}
