package main

import (
	"log"
	"os"
	routes "routes"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// reload .env at the first time
	err := godotenv.Load("../.env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}
func main() {
	// if APP_ENV == "production" then set gin mode to release mode
	if os.Getenv("APP_ENV") == "production" {
		gin.SetMode(gin.ReleaseMode)
	}

	// setting gin port
	port := os.Getenv("PORT")

	// create gin engine
	router := gin.Default()

	// load the all need template
	router.LoadHTMLGlob("../resources/views/**/*")

	// load all the routes
	routes.Routes(router)

	router.Run(":" + port)
}
