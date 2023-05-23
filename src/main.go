package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	routes "github.com/williamluisan/golang_auth0/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()
	
	routes.Routes(router)

	log.Fatal(router.Run(":4747"))
}