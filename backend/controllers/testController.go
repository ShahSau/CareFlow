package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTests(c *gin.Context) {

	fmt.Println("get all tests", c)

	c.JSON(200, gin.H{
		"message": "get all tests",
	})
}

func GetTest(c *gin.Context) {
	fmt.Println("get test", c)
}

func CreateTest(c *gin.Context) {
	fmt.Println("create test", c)
}

func UpdateTest(c *gin.Context) {
	fmt.Println("update test", c)
}

func DeleteTest(c *gin.Context) {
	fmt.Println("delete test", c)
}
