package handlers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

// User godoc
// @Summary get user details
// @Tags user
// @Produce  json
// @Security Bearer
// @Router /user/me [get]
func GetUserDetails(g *gin.Context) {
	fmt.Println("User details")
	g.JSON(200, gin.H{
		"message": "User details",
	})
}
