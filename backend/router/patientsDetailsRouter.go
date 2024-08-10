package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) PatientsDetailsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for patients details
	patientsDetailsRouteGrouping := rg.Group("/patients-details")
	patientsDetailsRouteGrouping.Use(cors.Default())

	// Get all patients details
	patientsDetailsRouteGrouping.GET("/", controllers.GetPatientsDetails)

	// Get a patient details
	patientsDetailsRouteGrouping.GET("/:id", controllers.GetPatientDetails)

	// Create a patient details
	patientsDetailsRouteGrouping.POST("/", controllers.CreatePatientDetails)

	// Update a patient details
	patientsDetailsRouteGrouping.PUT("/:id", controllers.UpdatePatientDetails)

	// Delete a patient details
	patientsDetailsRouteGrouping.DELETE("/:id", controllers.DeletePatientDetails)

}
