package middleware

import (
	"net/http"
	"strings"

	"github.com/Kimthean/go-chat/internal/repository"
	"github.com/Kimthean/go-chat/internal/utils"
	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		if authHeader == "" || tokenString == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Authorization header is missing or invalid"})
			return
		}

		claims, err := utils.ValidateToken(tokenString)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			return
		}

		user, err := repository.GetUserByID(claims.UserID.String())
		if err != nil || user == nil {
			errorMessage := "Failed to get user"
			if user == nil {
				errorMessage = "User not found"
			}

			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": errorMessage})
			return
		}

		c.Set("user", user)
		c.Set("userID", claims.UserID)
		c.Next()
	}
}
