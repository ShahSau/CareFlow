package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTestCategories(c *gin.Context) {

	fmt.Println("get all test categories", c)

	c.JSON(200, gin.H{
		"message": "get all test categories",
	})
}

func GetTestCategory(c *gin.Context) {
	fmt.Println("get test category", c)
}

func CreateTestCategory(c *gin.Context) {
	fmt.Println("create test category", c)
}

func UpdateTestCategory(c *gin.Context) {
	fmt.Println("update test category", c)
}

func DeleteTestCategory(c *gin.Context) {
	fmt.Println("delete test category", c)
}
