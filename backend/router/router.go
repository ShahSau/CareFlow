package router

import (
	"os"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

// ClientRoutes is a function that handles all the routes for the client side
func ClientRoutes() {
	r := routes{
		router: gin.Default(),
	}
	r.router.Use(gin.Logger())
	r.router.Use(gin.Recovery())

	// swagger TODO
	// r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// swagger docs TODO

	v := r.router.Group((os.Getenv("API_VERSION")))

	// all routes
	r.AuthRoutes(v)
	r.BuildingsRoutes(v)
	r.ConsultationRoutes(v)
	r.ConsultationsCateRoutes(v)
	r.ConsultationsDetailRoutes(v)
	r.DepartmentsRoutes(v)
	r.DieseasesRoutes(v)
	r.DocDetailsRoutes(v)
	r.DoctorsRoutes(v)
	r.ExpenseRoutes(v)
	r.MedicineRoutes(v)
	r.PatientRoutes(v)
	r.PatientsDetailsRoutes(v)
	r.RoomsRoutes(v)
	r.SpeacilizationRoutes(v)
	r.TestCateRoutes(v)
	r.TestRoutes(v)
	r.TreatmentRoutes(v)
	r.UserRoutes(v)
	r.VaccineAppointRoutes(v)
	r.VendorRoutes(v)

	r.router.Run(":" + os.Getenv("PORT"))
}
