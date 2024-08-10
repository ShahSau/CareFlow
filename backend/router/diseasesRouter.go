package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) DieseasesRoutes(rg *gin.RouterGroup) {
	// Create a new router group for diseases
	dieseasesRouteGrouping := rg.Group("/diseases")
	dieseasesRouteGrouping.Use(cors.Default())

	// Get all diseases
	dieseasesRouteGrouping.GET("/", controllers.GetDiseases)

	// Get a disease
	dieseasesRouteGrouping.GET("/:id", controllers.GetDisease)

	// Create a disease
	dieseasesRouteGrouping.POST("/", controllers.CreateDisease)

	// Update a disease
	dieseasesRouteGrouping.PUT("/:id", controllers.UpdateDisease)

	// Delete a disease
	dieseasesRouteGrouping.DELETE("/:id", controllers.DeleteDisease)

}
