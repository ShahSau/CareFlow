package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDepartments(c *gin.Context) {

	fmt.Println("get all departments", c)

	c.JSON(200, gin.H{
		"message": "get all departments",
	})
}

func GetDepartment(c *gin.Context) {
	fmt.Println("get department", c)
}

func CreateDepartment(c *gin.Context) {
	fmt.Println("create department", c)
}

func UpdateDepartment(c *gin.Context) {
	fmt.Println("update department", c)
}

func DeleteDepartment(c *gin.Context) {
	fmt.Println("delete department", c)
}
