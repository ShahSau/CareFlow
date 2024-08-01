package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetVendors(c *gin.Context) {

	fmt.Println("get all vendors", c)

	c.JSON(200, gin.H{
		"message": "get all vendors",
	})
}

func GetVendor(c *gin.Context) {
	fmt.Println("get vendor", c)
}

func CreateVendor(c *gin.Context) {
	fmt.Println("create vendor", c)
}

func UpdateVendor(c *gin.Context) {
	fmt.Println("update vendor", c)
}

func DeleteVendor(c *gin.Context) {
	fmt.Println("delete vendor", c)
}
