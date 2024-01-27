package main

import (
	"github.com/lowerzedo/appointment-booking-backend/initializers"
	"github.com/lowerzedo/appointment-booking-backend/models"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
}

func main() {
	initializers.DB.AutoMigrate(&models.Staff{})
}
