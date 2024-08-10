package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) DoctorsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for doctors
	doctorsRouteGrouping := rg.Group("/doctors")
	doctorsRouteGrouping.Use(cors.Default())

	// Get all doctors
	doctorsRouteGrouping.GET("/", controllers.GetDoctors)

	// Get a doctor
	doctorsRouteGrouping.GET("/:id", controllers.GetDoctor)

	// Create a doctor
	doctorsRouteGrouping.POST("/", controllers.CreateDoctor)

	// Update a doctor
	doctorsRouteGrouping.PUT("/:id", controllers.UpdateDoctor)

	// Delete a doctor
	doctorsRouteGrouping.DELETE("/:id", controllers.DeleteDoctor)

}
