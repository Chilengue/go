package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loadding .env file")
	}

	fmt.Println(os.Getenv(("TEST")))
}
