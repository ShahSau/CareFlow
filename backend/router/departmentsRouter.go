package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) DepartmentsRoutes(rg *gin.RouterGroup) {
	// Create a new router group for departments
	departmentsRouteGrouping := rg.Group("/departments")

	// Get all departments
	departmentsRouteGrouping.GET("/", controllers.GetDepartments)

	// Get a department
	departmentsRouteGrouping.GET("/:id", controllers.GetDepartment)

	// Create a department
	departmentsRouteGrouping.POST("/", controllers.CreateDepartment)

	// Update a department
	departmentsRouteGrouping.PUT("/:id", controllers.UpdateDepartment)

	// Delete a department
	departmentsRouteGrouping.DELETE("/:id", controllers.DeleteDepartment)

}
