package routes

import (
	"example.com/go-udemy-api/handlers"
	"example.com/go-udemy-api/middlewares"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {
	/* Events related Protected Routes */
	protected := server.Group("/")
	protected.Use(middlewares.Authenticate)
	protected.GET("/events", handlers.GetEvents)
	protected.GET("/events/:id", handlers.GetEventsById)
	protected.PUT("/events/:id", handlers.UpdateEvent)
	protected.DELETE("/events/:id", handlers.DeleteEvent)
	protected.POST("/events", handlers.SaveEvent)

	/* User Related */
	server.GET("/users", handlers.GetUsers)
	server.POST("/users/signup", handlers.CreateUser)
	server.POST("/users/login", handlers.Login)
	server.GET("/users/:id", handlers.GetUsers)

}
