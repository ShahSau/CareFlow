package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) PrescriptionsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for prescriptions
	prescriptionsRouteGrouping := rg.Group("/prescriptions")

	// Get all prescriptions
	prescriptionsRouteGrouping.GET("/", controllers.GetPrescriptions)

	// Get a prescription
	prescriptionsRouteGrouping.GET("/:id", controllers.GetPrescription)

	// Create a prescription
	prescriptionsRouteGrouping.POST("/", controllers.CreatePrescription)

	// Update a prescription
	prescriptionsRouteGrouping.PUT("/:id", controllers.UpdatePrescription)

	// Delete a prescription
	prescriptionsRouteGrouping.DELETE("/:id", controllers.DeletePrescription)

}
