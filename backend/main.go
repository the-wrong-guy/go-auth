package main

import (
	"fmt"

	"go-auth/auth"
	"go-auth/controllers"
	"go-auth/middleware"
	"go-auth/redis"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func init() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file:", err)
		return
	}
	// initialize redis client
	redis.InitRedis()
}

func main() {

	router := gin.Default()

	auth.InitAuth()

	router.Use(middleware.CORSMiddleware())
	router.GET("/auth/:provider/callback", controllers.HandleAuthCallback)
	router.GET("/auth/:provider", controllers.HandleAuthProvider)
	router.GET("/logout/:provider", controllers.LogoutHandler)
	router.Use(middleware.AuthMiddleware())
	router.GET("/check", controllers.CheckHandler)

	fmt.Println("Listening on localhost:3000")
	if err := router.Run(":3000"); err != nil {
		fmt.Println("Error starting server:", err)
	}
}
