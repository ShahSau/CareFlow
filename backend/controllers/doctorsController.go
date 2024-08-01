package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDoctors(c *gin.Context) {

	fmt.Println("get all doctors", c)

	c.JSON(200, gin.H{
		"message": "get all doctors",
	})
}

func GetDoctor(c *gin.Context) {
	fmt.Println("get doctor", c)
}

func CreateDoctor(c *gin.Context) {
	fmt.Println("create doctor", c)
}

func UpdateDoctor(c *gin.Context) {
	fmt.Println("update doctor", c)
}

func DeleteDoctor(c *gin.Context) {
	fmt.Println("delete doctor", c)
}
