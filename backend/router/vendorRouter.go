package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) VendorRoutes(rg *gin.RouterGroup) {
	// Create a new router group for vendors
	vendorRouteGrouping := rg.Group("/vendors")

	// Get all vendors
	vendorRouteGrouping.GET("/", controllers.GetVendors)

	// Get a vendor
	vendorRouteGrouping.GET("/:id", controllers.GetVendor)

	// Create a vendor
	vendorRouteGrouping.POST("/", controllers.CreateVendor)

	// Update a vendor
	vendorRouteGrouping.PUT("/:id", controllers.UpdateVendor)

	// Delete a vendor
	vendorRouteGrouping.DELETE("/:id", controllers.DeleteVendor)

}
