package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) ConsultationsDetailRoutes(rg *gin.RouterGroup) {
	// Create a new router group for consultationsDetail
	consultationsDetailRouteGrouping := rg.Group("/consultations-detail")

	// Get all consultationsDetail
	consultationsDetailRouteGrouping.GET("/", controllers.GetConsultationsDetail)

	// Get a consultationsDetail
	consultationsDetailRouteGrouping.GET("/:id", controllers.GetConsultationsDetail)

	// Create a consultationsDetail
	consultationsDetailRouteGrouping.POST("/", controllers.CreateConsultationsDetail)

	// Update a consultationsDetail
	consultationsDetailRouteGrouping.PUT("/:id", controllers.UpdateConsultationsDetail)

	// Delete a consultationsDetail
	consultationsDetailRouteGrouping.DELETE("/:id", controllers.DeleteConsultationsDetail)

}
