package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPatientsDetails(c *gin.Context) {

	fmt.Println("get all patients details", c)

	c.JSON(200, gin.H{
		"message": "get all patients details",
	})
}

func GetPatientDetails(c *gin.Context) {
	fmt.Println("get patient details", c)
}

func CreatePatientDetails(c *gin.Context) {
	fmt.Println("create patient details", c)
}

func UpdatePatientDetails(c *gin.Context) {
	fmt.Println("update patient details", c)
}

func DeletePatientDetails(c *gin.Context) {
	fmt.Println("delete patient details", c)
}
