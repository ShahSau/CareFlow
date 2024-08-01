package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetTreatments(c *gin.Context) {

	fmt.Println("get all treatments", c)

	c.JSON(200, gin.H{
		"message": "get all treatments",
	})
}

func GetTreatment(c *gin.Context) {
	fmt.Println("get treatment", c)
}

func CreateTreatment(c *gin.Context) {
	fmt.Println("create treatment", c)
}

func UpdateTreatment(c *gin.Context) {
	fmt.Println("update treatment", c)
}

func DeleteTreatment(c *gin.Context) {
	fmt.Println("delete treatment", c)
}

func GetTreatmentByPatient(c *gin.Context) {
	fmt.Println("get treatment by patient", c)
}

func GetTreatmentByCategory(c *gin.Context) {
	fmt.Println("get treatment by category", c)
}
