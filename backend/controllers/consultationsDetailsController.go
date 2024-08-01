package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetConsultationsDetails(c *gin.Context) {

	fmt.Println("get all consultationsDetails", c)

	c.JSON(200, gin.H{
		"message": "get all consultationsDetails",
	})
}

func GetConsultationsDetail(c *gin.Context) {
	fmt.Println("get consultationsDetail", c)
}

func CreateConsultationsDetail(c *gin.Context) {
	fmt.Println("create consultationsDetail", c)
}

func UpdateConsultationsDetail(c *gin.Context) {
	fmt.Println("update consultationsDetail", c)
}

func DeleteConsultationsDetail(c *gin.Context) {
	fmt.Println("delete consultationsDetail", c)
}
