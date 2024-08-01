package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetConsultationsCategories(c *gin.Context) {

	fmt.Println("get all consultations categories", c)

	c.JSON(200, gin.H{
		"message": "get all consultations categories",
	})
}

func GetConsultationsCategory(c *gin.Context) {
	fmt.Println("get consultations category", c)
}

func CreateConsultationsCategory(c *gin.Context) {
	fmt.Println("create consultations category", c)
}

func UpdateConsultationsCategory(c *gin.Context) {
	fmt.Println("update consultations category", c)
}

func DeleteConsultationsCategory(c *gin.Context) {
	fmt.Println("delete consultations category", c)
}
