package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) BuildingsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for buildings
	buildingsRouteGrouping := rg.Group("/buildings")
	buildingsRouteGrouping.Use(cors.Default())

	// Get all buildings
	buildingsRouteGrouping.GET("/", controllers.GetBuildings)

	// Get a building
	buildingsRouteGrouping.GET("/:id", controllers.GetBuilding)

	// Create a building
	buildingsRouteGrouping.POST("/", controllers.CreateBuilding)

	// Update a building
	buildingsRouteGrouping.PUT("/:id", controllers.UpdateBuilding)

	// Delete a building
	buildingsRouteGrouping.DELETE("/:id", controllers.DeleteBuilding)

}
