package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/lowerzedo/appointment-booking-backend/initializers"
	"github.com/lowerzedo/appointment-booking-backend/models"
)

func StaffCreate(c *gin.Context) {
	// Get data off req body
	var body struct {
		FullName string
		Role     string
		Status   bool
	}
	c.Bind(&body)

	// Create a staff
	staff := models.Staff{FullName: body.FullName, Role: body.Role, Status: body.Status}

	result := initializers.DB.Create(&staff)

	if result.Error != nil {
		c.Status(400)
		return
	}

	// Return it
	c.JSON(200, gin.H{
		"staff": staff,
	})
}

func StaffIndex(c *gin.Context) {
	// Get Staffs
	var staffs []models.Staff
	initializers.DB.Find(&staffs)

	// Responds with them
	c.JSON(200, gin.H{
		"staffs": staffs,
	})

}

func StaffById(c *gin.Context) {
	// Get staff id from url
	id := c.Param("id")

	// Fetch staff
	var staff models.Staff
	initializers.DB.First(&staff, id)

	// Respond with staff
	c.JSON(200, gin.H{
		"staff": staff,
	})
}

func StaffUpdate(c *gin.Context) {
	// Get the staff to be updated
	id := c.Param("id")

	// Get the data off req body
	var body struct {
		FullName string
		Role     string
		Status   bool
	}

	c.Bind(&body)

	// Find the staff we are updating
	var staff models.Staff
	initializers.DB.First(&staff, id)

	// Update it
	initializers.DB.Model(&staff).Updates(models.Staff{FullName: body.FullName, Role: body.Role, Status: body.Status})

	// Respond with it
	c.JSON(200, gin.H{
		"staff": staff,
	})
}

func StaffDelete(c *gin.Context) {
	// Get the id
	id := c.Param("id")

	// Delete the posts
	initializers.DB.Delete(&models.Staff{}, id)

	// Respond
	c.JSON(200, gin.H{
		"message": "Staff deleted!",
	})

}
