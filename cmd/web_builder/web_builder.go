package main

import (
	"net/http"
	"os"
	"strings"

	git_internal "github.com/DogAndHerDude/web-builder/git"
	"github.com/DogAndHerDude/web-builder/internal/app/auth/auth_handlers"
	"github.com/DogAndHerDude/web-builder/internal/app/auth/auth_service"
	"github.com/DogAndHerDude/web-builder/internal/app/db"
	"github.com/DogAndHerDude/web-builder/internal/app/site/site_handlers"
	"github.com/DogAndHerDude/web-builder/internal/app/site/site_service"
	"github.com/DogAndHerDude/web-builder/internal/app/user/user_handlers"
	"github.com/DogAndHerDude/web-builder/internal/app/user/user_service"
	"github.com/DogAndHerDude/web-builder/internal/pkg/builder"
	"github.com/DogAndHerDude/web-builder/internal/pkg/env"
	custom_middleware "github.com/DogAndHerDude/web-builder/middleware"
	"github.com/DogAndHerDude/web-builder/publisher"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/labstack/gommon/log"
)

func setupCORS(s *echo.Echo) {
	origin := os.Getenv("ORIGIN")
	allowed := strings.Split(origin, ",")

	s.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		Skipper:      middleware.DefaultSkipper,
		AllowOrigins: allowed,
		AllowMethods: []string{
			http.MethodGet,
			http.MethodHead,
			http.MethodPut,
			http.MethodPatch,
			http.MethodPost,
			http.MethodDelete,
		},
	}))
}

func main() {
	env.Init()

	DB := db.New()
	authService := auth_service.New()
	userService := user_service.New(DB)
	gitService := git_internal.New()
	builderService := builder.New()
	publisherService := publisher.New(gitService)
	siteService := site_service.New(DB, builderService, publisherService)

	server := echo.New()
	server.Validator = custom_middleware.NewValidator()
	apiGroup := server.Group("/api")

	setupCORS(server)
	server.Use(middleware.CSRF())
	server.Use(middleware.Logger())
	server.Logger.SetLevel(log.DEBUG)
	auth_handlers.RegisterHandlers(apiGroup, userService, authService)
	user_handlers.RegisterHandlers(userService, apiGroup)
	site_handlers.RegisterHandlers(apiGroup, siteService)

	err := server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		server.Logger.Fatal(err)
	}
}
