package handlers

import (
	"net/http"
	"strconv"

	"example.com/go-udemy-api/models"
	"github.com/gin-gonic/gin"
)

func GetEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error in getting events"})
		return
	}
	context.JSON(http.StatusOK, events)
}

func GetEventsById(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "id not correct"})
		return
	}

	event, err := models.GetEventById(id)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found"})
		return
	}
	context.JSON(http.StatusOK, event)
}

func SaveEvent(context *gin.Context) {
	var event models.Event
	err := context.ShouldBindJSON(&event)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	eventData, err := event.Save()
	if err != nil {
		context.JSON(http.StatusNotModified, gin.H{
			"error": "error in saving event",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{
		"message": "Event saved successfully",
		"event":   eventData,
	})
}

func UpdateEvent(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "id not correct"})
		return
	}

	var event models.Event
	err2 := context.ShouldBindJSON(&event)
	if err2 != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Wrong event data"})
		return
	}

	_, err3 := models.GetEventById(id)
	if err3 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found or DB error"})
		return
	}

	event.ID = id
	err4 := models.UpdateEvent(event)
	if err4 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not updated"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event updated Successfully"})
}

func DeleteEvent(context *gin.Context) {
	idStr := context.Param("id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "id not correct"})
		return
	}

	_, err3 := models.GetEventById(id)
	if err3 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not found or DB error"})
		return
	}

	err4 := models.DeleteEvent(id)
	if err4 != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": "Event not deleted"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Event deleted Successfully"})
}
