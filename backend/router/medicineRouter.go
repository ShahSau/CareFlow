package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) MedicineRoutes(rg *gin.RouterGroup) {
	// Create a new router group for medicines
	medicinesRouteGrouping := rg.Group("/medicines")

	// Get all medicines
	medicinesRouteGrouping.GET("/", controllers.GetMedicines)

	// Get a medicine
	medicinesRouteGrouping.GET("/:id", controllers.GetMedicine)

	// Create a medicine
	medicinesRouteGrouping.POST("/", controllers.CreateMedicine)

	// Update a medicine
	medicinesRouteGrouping.PUT("/:id", controllers.UpdateMedicine)

	// Delete a medicine
	medicinesRouteGrouping.DELETE("/:id", controllers.DeleteMedicine)

	// Get all medicines by disease
	medicinesRouteGrouping.GET("/disease/:id", controllers.GetMedicineByDisease)

	// Get all medicines by specialization
	medicinesRouteGrouping.GET("/specialization/:id", controllers.GetMedicineBySpecialization)

	// Get all medicines by prescription
	medicinesRouteGrouping.GET("/prescription/:id", controllers.GetMedicineByPrescription)

	// Get all medicines by patient
	medicinesRouteGrouping.GET("/patient/:id", controllers.GetMedicineByPatient)

}
