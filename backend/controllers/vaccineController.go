package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetVaccines(c *gin.Context) {

	fmt.Println("get all vaccines", c)

	c.JSON(200, gin.H{
		"message": "get all vaccines",
	})
}

func GetVaccine(c *gin.Context) {
	fmt.Println("get vaccine", c)
}

func CreateVaccine(c *gin.Context) {
	fmt.Println("create vaccine", c)
}

func UpdateVaccine(c *gin.Context) {
	fmt.Println("update vaccine", c)
}

func DeleteVaccine(c *gin.Context) {
	fmt.Println("delete vaccine", c)
}
