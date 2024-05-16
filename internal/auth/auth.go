package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Kimthean/go-chat/internal/models"
	"github.com/Kimthean/go-chat/internal/repository"
	"github.com/Kimthean/go-chat/internal/types"
	"github.com/Kimthean/go-chat/internal/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Response struct {
	Message string `json:"message"`
}

// SignUpHandler godoc
// @Summary Sign up a new user
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body types.SignUpRequest true "Sign up user"
// @Router /auth/signup [post]
func SignUpHandler(c *gin.Context) {
	var req types.SignUpRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, Response{Message: "Invalid request body"})
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, Response{Message: "Failed to hash password"})
		return
	}

	err = repository.CreateUser(req.Username, req.Email, hashedPassword)
	if err != nil {
		fmt.Println(err)
		if errors.Is(err, gorm.ErrDuplicatedKey) {
			c.JSON(http.StatusConflict, Response{Message: "Email already registered"})
		} else {
			c.JSON(http.StatusInternalServerError, Response{Message: "Failed to create user"})
		}
		return
	}

	c.JSON(http.StatusCreated, Response{Message: "User created"})
}

// LoginHandler godoc
// @Summary Log in a user
// @Description Log in a user with email and password
// @Tags auth
// @Accept  json
// @Produce  json
// @Param user body types.LoginRequest true "Login user"
// @Router /auth/login [post]
func LoginHandler(c *gin.Context) {
	var req types.LoginRequest

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Please sign up first"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	if !utils.ComparePassword(user.Password, []byte(req.Password)) {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid password"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(user.ID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	session := models.Session{
		UserID: user.ID,
		Token:  refreshToken,
	}

	if err := repository.CreateOrUpdateSession(&session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// RefreshTokenHandler godoc
// @Summary Refresh access token
// @Tags auth
// @Accept  json
// @Produce  json
// @Param token body types.RefreshTokenRequest true "Refresh token"
// @Router /auth/refresh [post]
func RefreshTokenHandler(c *gin.Context) {
	var req types.RefreshTokenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	fmt.Println(req.RefreshToken)

	session, err := repository.GetSessionByToken(req.RefreshToken)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	claims, err := utils.ValidateToken(req.RefreshToken)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid refresh token"})
		return
	}

	accessToken, refreshToken, err := utils.GenerateTokens(claims.UserID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate tokens"})
		return
	}

	session.Token = refreshToken
	if err := repository.CreateOrUpdateSession(session); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update session"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

// LogoutHandler godoc
// @Summary Log out a user
// @Tags auth
// @Security Bearer
// @Produce  json
// @Router /auth/logout [post]
func LogoutHandler(g *gin.Context) {

	userId, exist := g.Get("userID")
	if !exist {
		g.JSON(http.StatusUnauthorized, gin.H{"error": "User not authenticated"})
		return
	}

	session, err := repository.GetSessionById(userId.(uuid.UUID))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			g.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		} else {
			g.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
		}
		return
	}

	if err := repository.DeleteSession(session.Token); err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete session"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "Logged out"})
}
