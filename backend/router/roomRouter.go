package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) RoomsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for rooms
	roomsRouteGrouping := rg.Group("/rooms")
	roomsRouteGrouping.Use(cors.Default())

	// Get all rooms
	roomsRouteGrouping.GET("/", controllers.GetRooms)

	// Get a room
	roomsRouteGrouping.GET("/:id", controllers.GetRoom)

	// Create a room
	roomsRouteGrouping.POST("/", controllers.CreateRoom)

	// Update a room
	roomsRouteGrouping.PUT("/:id", controllers.UpdateRoom)

	// Delete a room
	roomsRouteGrouping.DELETE("/:id", controllers.DeleteRoom)

	// Swap a room
	roomsRouteGrouping.POST("/swap", controllers.SwapRoom)

}
