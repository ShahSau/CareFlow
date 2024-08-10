package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) VaccineRoutes(rg *gin.RouterGroup) {
	// Create a new router group for vaccines
	vaccineRouteGrouping := rg.Group("/vaccines")
	vaccineRouteGrouping.Use(cors.Default())

	// Get all vaccines
	vaccineRouteGrouping.GET("/", controllers.GetVaccines)

	// Get a vaccine
	vaccineRouteGrouping.GET("/:id", controllers.GetVaccine)

	// Create a vaccine
	vaccineRouteGrouping.POST("/", controllers.CreateVaccine)

	// Update a vaccine
	vaccineRouteGrouping.PUT("/:id", controllers.UpdateVaccine)

	// Delete a vaccine
	vaccineRouteGrouping.DELETE("/:id", controllers.DeleteVaccine)

}
