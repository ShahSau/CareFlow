package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPatients(c *gin.Context) {

	fmt.Println("get all patients", c)

	c.JSON(200, gin.H{
		"message": "get all patients",
	})
}

func GetPatient(c *gin.Context) {
	fmt.Println("get patient", c)
}

func CreatePatient(c *gin.Context) {
	fmt.Println("create patient", c)
}

func UpdatePatient(c *gin.Context) {
	fmt.Println("update patient", c)
}

func DeletePatient(c *gin.Context) {
	fmt.Println("delete patient", c)
}
