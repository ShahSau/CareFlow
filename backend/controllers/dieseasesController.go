package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDiseases(c *gin.Context) {

	fmt.Println("get all diseases", c)

	c.JSON(200, gin.H{
		"message": "get all diseases",
	})
}

func GetDisease(c *gin.Context) {
	fmt.Println("get disease", c)
}

func CreateDisease(c *gin.Context) {
	fmt.Println("create disease", c)
}

func UpdateDisease(c *gin.Context) {
	fmt.Println("update disease", c)
}

func DeleteDisease(c *gin.Context) {
	fmt.Println("delete disease", c)
}
