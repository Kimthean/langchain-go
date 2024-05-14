package handlers

import (
	"net/http"

	"github.com/Kimthean/go-chat/internal/repository"
	"github.com/Kimthean/go-chat/internal/types"
	"github.com/gin-gonic/gin"
)

// SignUpHandler godoc
// @Summary Sign up a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Router /signup [post]
func SignUpHandler(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.CreateUser(req.Username, req.Email, req.Password)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, user)
}

// LoginHandler godoc
// @Summary Log in a user
// @Description Log in a user with email and password
// @Tags users
// @Accept  json
// @Produce  json

// @Router /login [post]
func LoginHandler(c *gin.Context) {
	var req types.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := repository.GetUserByEmail(req.Email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}
