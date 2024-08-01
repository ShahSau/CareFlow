package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetSpecializations(c *gin.Context) {

	fmt.Println("get all specializations", c)

	c.JSON(200, gin.H{
		"message": "get all specializations",
	})
}

func GetSpecialization(c *gin.Context) {
	fmt.Println("get specialization", c)
}

func CreateSpecialization(c *gin.Context) {
	fmt.Println("create specialization", c)
}

func UpdateSpecialization(c *gin.Context) {
	fmt.Println("update specialization", c)
}

func DeleteSpecialization(c *gin.Context) {
	fmt.Println("delete specialization", c)
}
