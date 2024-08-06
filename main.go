package main

import (
	"log"

	"github.com/Chilengue/go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run("8000"); err != nil {
		log.Fatal(err)
	}
}
