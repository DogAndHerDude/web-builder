package auth

import (
	"net/http"

	hash_utils "github.com/DogAndHerDude/web-builder/internal/pkg/hash_utils"
	"github.com/DogAndHerDude/web-builder/user"

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
	userService user.IUserService
	authService IAuthService
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

	token, err := h.authService.GenerateJWT(ClaimValues{
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
	c.String(http.StatusAccepted, "Accepted")

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

func RegisterHandlers(e *echo.Group, u user.IUserService, a IAuthService) {
	h := &AuthHandlers{
		userService: u,
		authService: a,
	}
	subGroup := e.Group("/auth")
	subGroup.POST("/signup", h.Signup)
	subGroup.POST("/authenticate", h.Authenticate)
}
