package main

import (
	"github.com/croatiangrn/aoe-backend-challenge/app"
	"github.com/gin-gonic/gin"
	"log"
	"os"
)

const SuperHeroesJSONPath = "superheroes.json"

func test() {
}

func main() {
	app.Config = app.NewApplicationConfig(SuperHeroesJSONPath)

	if _, err := os.Stat(SuperHeroesJSONPath); err != nil {
		log.Fatalf("Error while opening superheroes.json: %s", err)
	}

	test()

	r := gin.Default()

	if err := r.Run(":8008"); err != nil {
		log.Fatalf("Couldn't run server: %s", err)
		return
	}
}
