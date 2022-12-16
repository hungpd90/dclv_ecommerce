package main

import (
	"fmt"
	"log"
	"os"

	"backend/routes"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	// os.Setenv("TWILIO_ACCOUNT_SID", "")
	// os.Setenv("TWILIO_TOKEN", "")
	// os.Setenv("SECRET_LOVE", "")
	if port == "" {
		port = "8000"
	}

	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Logger())
	routes.Routes(router)
	fmt.Println(os.Getenv("TWILIO_ACCOUNT_SID"))
	fmt.Println(os.Getenv("TWILIO_TOKEN"))
	fmt.Println(os.Getenv("TWILIO_FROM_PHONE_NUMBER"))
	log.Fatal(router.Run(":" + port))
}
