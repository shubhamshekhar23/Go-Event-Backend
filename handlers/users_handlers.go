package handlers

import (
	"net/http"

	"example.com/go-udemy-api/models"
	"example.com/go-udemy-api/utils"
	"github.com/gin-gonic/gin"
)

func GetUsers(context *gin.Context) {
	users, err := models.GetAllUsers()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting users"})
		return
	}
	context.JSON(http.StatusOK, users)
}

// func GetUserById(context *gin.Context) {
// 	idStr := context.Param("id")
// 	id, err := strconv.ParseInt(idStr, 10, 64)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "id not correct"})
// 		return
// 	}

// 	event, err := models.GetEventById(id)
// 	if err != nil {
// 		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
// 		return
// 	}
// 	context.JSON(http.StatusOK, event)
// }

func CreateUser(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, err := user.Save()
	if err != nil {
		context.JSON(http.StatusNotModified, gin.H{
			"error": "error in creating user",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User created successfully",
		"user":    userData,
	})
}

func Login(context *gin.Context) {
	var user models.User
	err := context.ShouldBindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userData, err := user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"error": err.Error(),
		})
		return
	}

	token, err2 := utils.GenerateToken(userData.Email, userData.ID)
	if err2 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": "Something Bad Happened",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "User Loggedin successfully",
		"token":   token,
	})
}
