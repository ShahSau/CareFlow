package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) AuthRoutes(rg *gin.RouterGroup) {
	// Create a new router group for auth
	authRouteGrouping := rg.Group("/auth")

	// Login
	authRouteGrouping.POST("/login", controllers.Login)

	// Register
	authRouteGrouping.POST("/register", controllers.Register)

	// Logout
	authRouteGrouping.POST("/logout", controllers.Logout)

}
