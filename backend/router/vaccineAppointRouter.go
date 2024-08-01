package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) VaccineAppointRoutes(rg *gin.RouterGroup) {
	// Create a new router group for vaccine appointments
	vaccineAppointRouteGrouping := rg.Group("/vaccine-appointments")

	// Get all vaccine appointments
	vaccineAppointRouteGrouping.GET("/", controllers.GetVaccineAppointments)

	// Get a vaccine appointment
	vaccineAppointRouteGrouping.GET("/:id", controllers.GetVaccineAppointment)

	// Create a vaccine appointment
	vaccineAppointRouteGrouping.POST("/", controllers.CreateVaccineAppointment)

	// Update a vaccine appointment
	vaccineAppointRouteGrouping.PUT("/:id", controllers.UpdateVaccineAppointment)

	// Delete a vaccine appointment
	vaccineAppointRouteGrouping.DELETE("/:id", controllers.DeleteVaccineAppointment)

	// Get vaccine appointments by vaccine
	vaccineAppointRouteGrouping.GET("/vaccine/:vaccineId", controllers.GetVaccineAppointmentsByVaccine)

	// Get vaccine appointments by user
	vaccineAppointRouteGrouping.GET("/user/:userId", controllers.GetVaccineAppointmentsByUser)

}
