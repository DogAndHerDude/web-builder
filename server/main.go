package main

import (
	"app/db"
	"app/env"
)

func main() {
	env.Init()
	db.Init()
}
