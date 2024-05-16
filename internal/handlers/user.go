package handlers

import (
	"github.com/gin-gonic/gin"
)

// User godoc
// @Summary get user details
// @Tags user
// @Produce  json
// @Security Bearer
// @Router /user/me [get]
func GetUserDetails(g *gin.Context) {
	user, exist := g.Get("user")
	if !exist {
		g.JSON(500, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	g.JSON(200, gin.H{
		"data": user,
	})
}
