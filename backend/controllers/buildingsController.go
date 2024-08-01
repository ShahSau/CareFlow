package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetBuildings(c *gin.Context) {

	fmt.Println("get all buildings", c)

	c.JSON(200, gin.H{
		"message": "get all buildings",
	})
}

func GetBuilding(c *gin.Context) {
	fmt.Println("get building", c)
}

func CreateBuilding(c *gin.Context) {
	fmt.Println("create building", c)
}

func UpdateBuilding(c *gin.Context) {
	fmt.Println("update building", c)
}

func DeleteBuilding(c *gin.Context) {
	fmt.Println("delete building", c)
}
