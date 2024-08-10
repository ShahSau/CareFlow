package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) PatientRoutes(rg *gin.RouterGroup) {
	// Create a new router group for patients
	patientsRouteGrouping := rg.Group("/patients")
	patientsRouteGrouping.Use(cors.Default())

	// Get all patients
	patientsRouteGrouping.GET("/", controllers.GetPatients)

	// Get a patient
	patientsRouteGrouping.GET("/:id", controllers.GetPatient)

	// Create a patient
	patientsRouteGrouping.POST("/", controllers.CreatePatient)

	// Update a patient
	patientsRouteGrouping.PUT("/:id", controllers.UpdatePatient)

	// Delete a patient
	patientsRouteGrouping.DELETE("/:id", controllers.DeletePatient)

}
