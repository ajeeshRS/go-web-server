package main

import (
	"log"
	"os"
	"web-service/routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading env")
	}

	router := gin.Default()
	routes.SetupRouter(router)

	PORT := os.Getenv("PORT")
	router.Run(":" + PORT)
}
