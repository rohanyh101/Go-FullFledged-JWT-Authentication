package main

import (
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/rohanhonnakatti/golang-jwt-auth/routes"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("error loading .env file: %v", err)
	}

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}

	router := gin.New()
	router.Use(gin.Logger())

	routes.AuthRoutes(router)

	// the below endpoint will get through, bcs we are not using authMiddleware till now
	router.GET("/api-1", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for API-1"})
	})

	// the authMiddleware will get be implemented in the below router,
	// all endpoins comes after this must go through authentication layer...
	routes.UserRoutes(router)

	router.GET("/api-2", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"success": "Access granted for API-2"})
	})

	router.Run(":" + PORT)
}
