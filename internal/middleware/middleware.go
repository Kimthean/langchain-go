package middleware

import (
    "net/http"
    "strings"

    "github.com/Kimthean/go-chat/internal/utils"
    "github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
    return func(c *gin.Context) {
        authHeader := c.GetHeader("Authorization")
        if authHeader == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        if tokenString == "" {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        claims, err := utils.ValidateToken(tokenString)
        if err != nil {
            c.AbortWithStatus(http.StatusUnauthorized)
            return
        }

        c.Set("userID", claims.UserID)
        c.Next()
    }
}