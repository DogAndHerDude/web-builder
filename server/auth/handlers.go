package auth

import (
	"net/http"

	"app/user"
	"app/utils"

	"github.com/labstack/echo/v4"
)

// Needs dat avalidation
type SignupPayload struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignupResponse struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}

type AuthHandlers struct {
	userService user.IUserService
	authService IAuthService
}

func (h *AuthHandlers) Signup(c echo.Context) error {
	var payload SignupPayload
	err := c.Bind(&payload)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	salt, err := utils.RandomSecret(10)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	hashedPassword, err := utils.HashString(payload.Password, salt)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, "Something went wrong")
	}

	user, err := h.userService.CreateUser(payload.Email, string(salt), string(hashedPassword))
	if err != nil {
		return echo.NewHTTPError(http.StatusExpectationFailed)
	}
	token, err := h.authService.GenerateJWT(ClaimValues{
		ID: user.ID,
	})
	if err != nil {
		return echo.NewHTTPError(http.StatusUnprocessableEntity)
	}

	c.SetCookie(&http.Cookie{
		Name:     "Authorization",
		Value:    token,
		HttpOnly: true,
		SameSite: http.SameSiteStrictMode,
	})

	// c.JSON(http.StatusCreated)

	return nil
}

func (h *AuthHandlers) Authenticate(c echo.Context) error {
	return nil
}

func RegisterHandlers(e *echo.Group, u user.IUserService, a IAuthService) {
	h := &AuthHandlers{
		userService: u,
		authService: a,
	}
	subGroup := e.Group("/auth")
	subGroup.POST("/signup", h.Signup)
	subGroup.POST("/authenticate", h.Authenticate)
}
