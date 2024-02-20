package main

import (
	"app/builder"
	"app/db"
	"app/env"
	git_internal "app/git"
	"app/publisher"
	"app/site"
	"app/user"
)

func main() {
	env.Init()

	DB := db.New()
	userService := user.New(DB)
	gitService := git_internal.New()
	builderService := builder.New()
	publisherService := publisher.New(gitService)
	siteService := site.New(DB, builderService, publisherService)
}
