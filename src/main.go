package main

import (
	espressController "EspressoTracking/controllers/espresso"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	app := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"GET", "POST", "PUT", "DELETE"}
	config.AllowCredentials = true
	config.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}
	config.AllowWildcard = true
	app.Use(cors.New(config))

	espressController.EspressoApplicationRoutes("/api/espresso", app)

	//espressController.EspressoApplicationRoutes("/api/espresso", app)

	serverErr := app.Run(":3001")
	if serverErr != nil {
		panic(serverErr)
	}
}
