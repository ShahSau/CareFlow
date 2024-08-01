package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetRooms(c *gin.Context) {

	fmt.Println("get all rooms", c)

	c.JSON(200, gin.H{
		"message": "get all rooms",
	})
}

func GetRoom(c *gin.Context) {
	fmt.Println("get room", c)
}

func CreateRoom(c *gin.Context) {
	fmt.Println("create room", c)
}

func UpdateRoom(c *gin.Context) {
	fmt.Println("update room", c)
}

func DeleteRoom(c *gin.Context) {
	fmt.Println("delete room", c)
}

func SwapRoom(c *gin.Context) {
	fmt.Println("swap room", c)
}
