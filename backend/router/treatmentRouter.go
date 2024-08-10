package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) TreatmentRoutes(rg *gin.RouterGroup) {
	// Create a new router group for treatments
	treatmentRouteGrouping := rg.Group("/treatments")
	treatmentRouteGrouping.Use(cors.Default())

	// Get all treatments
	treatmentRouteGrouping.GET("/", controllers.GetTreatments)

	// Get a treatment
	treatmentRouteGrouping.GET("/:id", controllers.GetTreatment)

	// Create a treatment
	treatmentRouteGrouping.POST("/", controllers.CreateTreatment)

	// Update a treatment
	treatmentRouteGrouping.PUT("/:id", controllers.UpdateTreatment)

	// Delete a treatment
	treatmentRouteGrouping.DELETE("/:id", controllers.DeleteTreatment)

	// Get treatment by patient
	treatmentRouteGrouping.GET("/patient/:patientId", controllers.GetTreatmentByPatient)

	// Get treatment by category
	treatmentRouteGrouping.GET("/category/:categoryId", controllers.GetTreatmentByCategory)

}
