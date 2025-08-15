package handlers

import (
	"net/http"
	"strconv"

	"example.com/go-udemy-api/models"
	"github.com/gin-gonic/gin"
)

func RegisterEvent(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var register models.Register
	register.UserID = userId
	register.EventID = eventId

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	savedRegister, err := register.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message":  "Registration successful",
		"register": savedRegister,
	})
}

func DeleteRegistration(context *gin.Context) {
	userId := context.GetInt64("userId")
	eventIdStr := context.Param("id")
	eventId, err := strconv.ParseInt(eventIdStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Invalid event ID"})
		return
	}

	var register models.Register
	register.UserID = userId
	register.EventID = eventId

	_, err = models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": "Event not found"})
		return
	}

	err = register.Delete()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Delete Registration successful",
	})
}
