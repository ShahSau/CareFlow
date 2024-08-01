package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) TestRoutes(rg *gin.RouterGroup) {
	// Create a new router group for tests
	testRouteGrouping := rg.Group("/tests")

	// Get all tests
	testRouteGrouping.GET("/", controllers.GetTests)

	// Get a test
	testRouteGrouping.GET("/:id", controllers.GetTest)

	// Create a test
	testRouteGrouping.POST("/", controllers.CreateTest)

	// Update a test
	testRouteGrouping.PUT("/:id", controllers.UpdateTest)

	// Delete a test
	testRouteGrouping.DELETE("/:id", controllers.DeleteTest)

}
