package main

import (
	"example.com/go-udemy-api/db"
	"example.com/go-udemy-api/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()
	/* Register routes */
	routes.RegisterRoutes(server)
	server.Run(":8080")
}
