package auth

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/Kimthean/go-chat/internal/repository"
	"github.com/Kimthean/go-chat/internal/types"
	"github.com/Kimthean/go-chat/internal/utils"
	"github.com/gin-gonic/gin"
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
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
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
// @Failure 400 {object} types.ErrorResponse
// @Failure 500 {object} types.ErrorResponse
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
			c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
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

	c.JSON(http.StatusOK, gin.H{
		"access_token":  accessToken,
		"refresh_token": refreshToken,
	})
}

