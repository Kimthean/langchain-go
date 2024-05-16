package main

import (
	"log"

	docs "github.com/Kimthean/go-chat/docs"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"github.com/Kimthean/go-chat/internal/auth"
	"github.com/Kimthean/go-chat/internal/database"
	"github.com/Kimthean/go-chat/internal/handlers"
	"github.com/Kimthean/go-chat/internal/middleware"
	"github.com/gin-gonic/gin"
)

// @title           LangChain Go RAG Agent API
// @version         1.0
// @description     Langchain go RAG agent API.
// @termsOfService  http://swagger.io/terms/
// @host            localhost:8080
// @BasePath        /api/v1
// @securityDefinitions.apiKey Bearer
// @in header
// @name Authorization
// @scheme Bearer
func main() {
	if err := database.Connect(); err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	// db := database.DB

	// db.Migrator().DropTable(&models.User{}, &models.Session{})
	// db.AutoMigrate(&models.User{}, &models.Session{})

	r := gin.Default()
	docs.SwaggerInfo.Title = "Go Chat"
	docs.SwaggerInfo.Description = "A simple chat application"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.BasePath = "/api/v1"

	v1 := r.Group("/api/v1")
	{
		auththentication := v1.Group("/auth")
		{
			auththentication.POST("/signup", auth.SignUpHandler)
			auththentication.POST("/login", auth.LoginHandler)
			auththentication.POST("/refresh", auth.RefreshTokenHandler)
			auththentication.POST("/logout", middleware.AuthMiddleware(), auth.LogoutHandler)
		}
		user := v1.Group("/user", middleware.AuthMiddleware())
		{
			user.GET("/me", handlers.GetUserDetails)
			user.POST("/profile", handlers.UploadProfileImage)
		}
	}
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler, ginSwagger.PersistAuthorization(true)))

	r.Run(":8080")
}
