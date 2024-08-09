package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) AuthRoutes(rg *gin.RouterGroup) {
	// Create a new router group for auth
	authRouteGrouping := rg.Group("/auth")
	authRouteGrouping.Use(cors.Default())

	// Login
	authRouteGrouping.POST("/login", controllers.Login)

	// Register
	authRouteGrouping.POST("/register", controllers.Register)

	// Logout
	authRouteGrouping.POST("/logout", controllers.Logout)

	//test
	authRouteGrouping.GET("/test", controllers.Test)

}
