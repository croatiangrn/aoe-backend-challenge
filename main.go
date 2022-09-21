package main

import (
	"github.com/gin-gonic/gin"
	"log"
)

func main() {
	r := gin.Default()

	err := r.Run(":8008")

	if err != nil {
		log.Fatalln("Couldn't run server")
		return
	}
}
