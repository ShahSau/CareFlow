package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetPrescriptions(c *gin.Context) {

	fmt.Println("get all prescriptions", c)

	c.JSON(200, gin.H{
		"message": "get all prescriptions",
	})
}

func GetPrescription(c *gin.Context) {
	fmt.Println("get prescription", c)
}

func CreatePrescription(c *gin.Context) {
	fmt.Println("create prescription", c)
}

func UpdatePrescription(c *gin.Context) {
	fmt.Println("update prescription", c)
}

func DeletePrescription(c *gin.Context) {
	fmt.Println("delete prescription", c)
}
