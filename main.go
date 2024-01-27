package main

import (
	"github.com/gin-gonic/gin"
	"github.com/lowerzedo/appointment-booking-backend/controllers"
	"github.com/lowerzedo/appointment-booking-backend/initializers"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()
	r.POST("/staff", controllers.StaffCreate)
	r.PUT("/staff/:id", controllers.StaffUpdate)
	r.GET("/staff", controllers.StaffIndex)
	r.GET("/staff/:id", controllers.StaffById)
	r.DELETE("/staff/:id", controllers.StaffDelete)
	r.Run() // listen and serve on 0.0.0.0:8080
}
