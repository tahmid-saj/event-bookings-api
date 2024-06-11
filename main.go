package main

import (
	"tahmid-saj/event-booking-api/routes"

	// "tahmid-saj/event-booking-api/db"
	"github.com/gin-gonic/gin"
)

func main() {
	// db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080")
}