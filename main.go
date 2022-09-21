package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

func test() {

	fmt.Println("")
}

func main() {
	test()
	r := gin.Default()

	err := r.Run(":8008")

	if err != nil {
		log.Fatalln("Couldn't run server")
		return
	}
}
