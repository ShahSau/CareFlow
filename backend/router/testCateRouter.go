package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) TestCateRoutes(rg *gin.RouterGroup) {
	// Create a new router group for test categories
	testCateRouteGrouping := rg.Group("/testCategories")

	// Get all test categories
	testCateRouteGrouping.GET("/", controllers.GetTestCategories)

	// Get a test category
	testCateRouteGrouping.GET("/:id", controllers.GetTestCategory)

	// Create a test category
	testCateRouteGrouping.POST("/", controllers.CreateTestCategory)

	// Update a test category
	testCateRouteGrouping.PUT("/:id", controllers.UpdateTestCategory)

	// Delete a test category
	testCateRouteGrouping.DELETE("/:id", controllers.DeleteTestCategory)

}
