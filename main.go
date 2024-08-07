package main

import (
	"log"

	"github.com/Chilengue/go.git/src/controller/routes"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

// func main() {
// 	err := godotenv.Load()
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	router := gin.Default()

// 	routes.InitRoutes(&router.RouterGroup)

// 	if err := router.Run("8087"); err != nil {
// 		log.Fatal(err)
// 	}
// }

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loafing  .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8100"); err != nil {
		log.Fatal(err)
	}
}
