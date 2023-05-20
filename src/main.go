package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	
	router.Routes(router)

	log.Fatal(router.Run(":8080"))
}