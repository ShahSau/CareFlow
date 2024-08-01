package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetExpenses(c *gin.Context) {

	fmt.Println("get all expenses", c)

	c.JSON(200, gin.H{
		"message": "get all expenses",
	})
}

func GetExpense(c *gin.Context) {
	fmt.Println("get expense", c)
}

func CreateExpense(c *gin.Context) {
	fmt.Println("create expense", c)
}

func UpdateExpense(c *gin.Context) {
	fmt.Println("update expense", c)
}

func DeleteExpense(c *gin.Context) {
	fmt.Println("delete expense", c)
}

func GetExpenseByDepartment(c *gin.Context) {
	fmt.Println("get expense by department", c)
}
