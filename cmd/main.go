package main

import (
	"log"

	docs "github.com/Kimthean/go-chat/cmd/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Kimthean/go-chat/internal/database"
	"github.com/Kimthean/go-chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	r := gin.Default()
	docs.SwaggerInfo.Title = "Go Chat"
	docs.SwaggerInfo.Description = "A simple chat application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("/signup", handlers.SignUpHandler)
			auth.POST("/login", handlers.LoginHandler)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(":8080")
}
