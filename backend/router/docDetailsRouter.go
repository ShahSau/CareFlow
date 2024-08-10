package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) DocDetailsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for docDetails
	docDetailsRouteGrouping := rg.Group("/doc-details")
	docDetailsRouteGrouping.Use(cors.Default())

	// Get all docDetails
	docDetailsRouteGrouping.GET("/", controllers.GetDocDetails)

	// Get a docDetails
	docDetailsRouteGrouping.GET("/:id", controllers.GetDocDetail)

	// Create a docDetails
	docDetailsRouteGrouping.POST("/", controllers.CreateDocDetail)

	// Update a docDetails
	docDetailsRouteGrouping.PUT("/:id", controllers.UpdateDocDetail)

	// Delete a docDetails
	docDetailsRouteGrouping.DELETE("/:id", controllers.DeleteDocDetail)

}
