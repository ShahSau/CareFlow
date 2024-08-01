package router

import (
	"github.com/ShahSau/CareFlow/backend/controllers"
	"github.com/gin-gonic/gin"
)

func (r routes) ExpenseRoutes(rg *gin.RouterGroup) {
	// Create a new router group for expenses
	expensesRouteGrouping := rg.Group("/expenses")

	// Get all expenses
	expensesRouteGrouping.GET("/", controllers.GetExpenses)

	// Get a expense
	expensesRouteGrouping.GET("/:id", controllers.GetExpense)

	// Create a expense
	expensesRouteGrouping.POST("/", controllers.CreateExpense)

	// Update a expense
	expensesRouteGrouping.PUT("/:id", controllers.UpdateExpense)

	// Delete a expense
	expensesRouteGrouping.DELETE("/:id", controllers.DeleteExpense)

	// Get expense by department
	expensesRouteGrouping.GET("/department/:id", controllers.GetExpenseByDepartment)

}
