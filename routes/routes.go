package routes

import (
	"example.com/go-udemy-api/handlers"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	/* Events related */
	server.GET("/events", handlers.GetEvents)
	server.GET("/events/:id", handlers.GetEventsById)
	server.PUT("/events/:id", handlers.UpdateEvent)
	server.DELETE("/events/:id", handlers.DeleteEvent)
	server.POST("/events", handlers.SaveEvent)

	/* User Related */
	server.GET("/users", handlers.GetUsers)
	server.GET("/users/:id", handlers.GetUsers)
	server.POST("/users/signup", handlers.CreateUser)
	server.POST("/users/login", handlers.CreateUser)

}
