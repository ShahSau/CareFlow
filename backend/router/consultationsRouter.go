package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func (r routes) ConsultationRoutes(rg *gin.RouterGroup) {
	// Create a new router group for consultations
	consultationRouteGrouping := rg.Group("/consultations")
	consultationRouteGrouping.Use(cors.Default())

	// Get all consultations
	consultationRouteGrouping.GET("/", controllers.GetConsultations)

	// Get a consultation
	consultationRouteGrouping.GET("/:id", controllers.GetConsultation)

	// Create a consultation
	consultationRouteGrouping.POST("/", controllers.CreateConsultation)

	// Update a consultation
	consultationRouteGrouping.PUT("/:id", controllers.UpdateConsultation)

	// Delete a consultation
	consultationRouteGrouping.DELETE("/:id", controllers.DeleteConsultation)

	// Get consultation by patient
	consultationRouteGrouping.GET("/patient/:patientId", controllers.GetConsultationByPatient)

	// Get consultation by doctor
	consultationRouteGrouping.GET("/doctor/:doctorId", controllers.GetConsultationByDoctor)

}
