package main

import (
	"os"

	"app/auth"
	"app/builder"
	"app/db"
	"app/env"
	git_internal "app/git"
	"app/publisher"
	"app/site"
	"app/user"

	"github.com/labstack/echo/v4"
	"github.com/labstack/gommon/log"
)

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
	apiGroup := server.Group("/api")

	server.Logger.SetLevel(log.DEBUG)
	auth.RegisterHandlers(apiGroup, userService, authService)
	user.RegisterHandlers(userService, apiGroup)
	site.RegisterHandlers(apiGroup, siteService)

	err := server.Start(":" + os.Getenv("PORT"))
	if err != nil {
		server.Logger.Fatal(err)
	}
}
