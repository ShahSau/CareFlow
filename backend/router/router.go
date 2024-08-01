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
	r.DieseasesRoutes(v)
	r.PatientRoutes(v)
	r.RoomsRoutes(v)
	r.UserRoutes(v)
	r.VendorRoutes(v)
	r.VaccineRoutes(v)
	r.VaccineAppointRoutes(v)
	r.BuildingsRoutes(v)
	r.ConsultationRoutes(v)
	r.DoctorsRoutes(v)
	r.TestRoutes(v)
	r.TestCateRoutes(v)
	r.DocDetailsRoutes(v)
	r.DepartmentsRoutes(v)
	r.SpeacilizationRoutes(v)
	r.MedicineRoutes(v)
	r.ExpenseRoutes(v)

	r.router.Run(":" + os.Getenv("PORT"))
}
