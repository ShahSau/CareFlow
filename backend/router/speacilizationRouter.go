package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) SpeacilizationRoutes(rg *gin.RouterGroup) {
	// Create a new router group for speacilizations
	speacilizationRouteGrouping := rg.Group("/speacilizations")

	// Get all speacilizations
	speacilizationRouteGrouping.GET("/", controllers.GetSpecializations)

	// Get a speacilization
	speacilizationRouteGrouping.GET("/:id", controllers.GetSpecialization)

	// Create a speacilization
	speacilizationRouteGrouping.POST("/", controllers.CreateSpecialization)

	// Update a speacilization
	speacilizationRouteGrouping.PUT("/:id", controllers.UpdateSpecialization)

	// Delete a speacilization
	speacilizationRouteGrouping.DELETE("/:id", controllers.DeleteSpecialization)

}
