package main

import (
	"net/http"
	"os"
	"strings"

	"app/auth"
	"app/builder"
	"app/db"
	"app/env"
	git_internal "app/git"
	custom_middleware "app/middleware"
	"app/publisher"
	"app/site"
	"app/user"

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
	siteService := site.New(DB, builderService, publisherService)

	server := echo.New()
	server.Validator = custom_middleware.NewValidator()
	apiGroup := server.Group("/api")

	setupCORS(server)
	server.Use(middleware.Logger())
	server.Logger.SetLevel(log.DEBUG)
	auth.RegisterHandlers(apiGroup, userService, authService)
	user.RegisterHandlers(userService, apiGroup)
	site.RegisterHandlers(apiGroup, siteService)

	err := server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		server.Logger.Fatal(err)
	}
}
