package main

import (
	"app/appointment-booking/src/routes" // Assuming this is the correct import path for your routes package

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup routes
	routes.SetupRoutes(router) // Ensure that this function exists in the routes package

	// Start server
	router.Run(":8080")
}
