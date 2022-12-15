package main

import (
	"log"
	"os"

	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	os.Setenv("TWILIO_ACCOUNT_SID", "AC516676b8505af5a3d751e08a8a2c1278")
	os.Setenv("TWILIO_TOKEN", "8884850eadd72377a030107219ec16f9")
	os.Setenv("SECRET_LOVE", "")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	routes.Routes(router)

	log.Fatal(router.Run(":" + port))
}
