package routes

import (
	"github.com/Kimthean/go-chat/internal/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.RouterGroup) {
	r.POST("/signup", handlers.SignUpHandler)
	r.POST("/login", handlers.LoginHandler)
}
