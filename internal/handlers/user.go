package handlers

import (
	"fmt"
	"net/http"

	"github.com/Kimthean/go-chat/internal/utils"
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
	userId, _ := g.Get("userID")
	if !exist {
		g.JSON(500, gin.H{
			"error": "Failed to get user",
		})
		return
	}

	g.JSON(200, gin.H{
		"data":   user,
		"userId": userId,
	})
}

// UploadProfileImage godoc
// @Summary upload profile image
// @Tags user
// @Produce  json
// @Security Bearer
// @Router /user/profile [post]
// @Param file formData file true "Profile Image"
func UploadProfileImage(g *gin.Context) {
	// userId, _ := g.Get("userID")
	// user, err := repository.GetUserByID(userId.(string))
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// if user.ProfileImage != "" {
	// 	err := utils.DeleteFileFromS3(user.ProfileImage)
	// 	if err != nil {
	// 		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete file"})
	// 		return
	// 	}

	// }

	file, err := g.FormFile("file")
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{"error": "No file uploaded"})
		return
	}

	err = utils.UploadFileToS3(file.Filename, "profile")
	fmt.Println(err)
	if err != nil {
		g.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to upload file"})
		return
	}

	g.JSON(http.StatusOK, gin.H{"message": "File uploaded successfully"})
}
