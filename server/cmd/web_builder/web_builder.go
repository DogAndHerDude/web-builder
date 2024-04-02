package main

import (
	"net/http"
	"os"
	"strings"

	"github.com/DogAndHerDude/web-builder/auth"
	"github.com/DogAndHerDude/web-builder/builder"
	"github.com/DogAndHerDude/web-builder/db"
	"github.com/DogAndHerDude/web-builder/env"
	git_internal "github.com/DogAndHerDude/web-builder/git"
	"github.com/DogAndHerDude/web-builder/internal/app/site/site_handlers"
	"github.com/DogAndHerDude/web-builder/internal/app/site/site_service"
	custom_middleware "github.com/DogAndHerDude/web-builder/middleware"
	"github.com/DogAndHerDude/web-builder/publisher"
	"github.com/DogAndHerDude/web-builder/user"

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
	authService := auth.New()
	userService := user.New(DB)
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
	auth.RegisterHandlers(apiGroup, userService, authService)
	user.RegisterHandlers(userService, apiGroup)
	site_handlers.RegisterHandlers(apiGroup, siteService)

	err := server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		server.Logger.Fatal(err)
	}
}
