package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) UserRoutes(rg *gin.RouterGroup) {
	// Create a new router group for users
	usersRouteGrouping := rg.Group("/users")

	// Get all users
	usersRouteGrouping.GET("/", controllers.GetUsers)

	// Get a user
	usersRouteGrouping.GET("/:id", controllers.GetUser)

	// Create a user
	usersRouteGrouping.POST("/", controllers.CreateUser)

	// Update a user
	usersRouteGrouping.PUT("/:id", controllers.UpdateUser)

	// Delete a user
	usersRouteGrouping.DELETE("/:id", controllers.DeleteUser)

}
