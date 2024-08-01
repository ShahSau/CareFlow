package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) ConsultationsCateRoutes(rg *gin.RouterGroup) {
	// Create a new router group for consultations categories
	consultationsCateRouteGrouping := rg.Group("/consultations-categories")

	// Get all consultations categories
	consultationsCateRouteGrouping.GET("/", controllers.GetConsultationsCategories)

	// Get a consultations category
	consultationsCateRouteGrouping.GET("/:id", controllers.GetConsultationsCategory)

	// Create a consultations category
	consultationsCateRouteGrouping.POST("/", controllers.CreateConsultationsCategory)

	// Update a consultations category
	consultationsCateRouteGrouping.PUT("/:id", controllers.UpdateConsultationsCategory)

	// Delete a consultations category
	consultationsCateRouteGrouping.DELETE("/:id", controllers.DeleteConsultationsCategory)

}
