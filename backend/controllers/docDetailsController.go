package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetDocDetails(c *gin.Context) {

	fmt.Println("get all docDetails", c)

	c.JSON(200, gin.H{
		"message": "get all docDetails",
	})
}

func GetDocDetail(c *gin.Context) {
	fmt.Println("get docDetail", c)
}

func CreateDocDetail(c *gin.Context) {
	fmt.Println("create docDetail", c)
}

func UpdateDocDetail(c *gin.Context) {
	fmt.Println("update docDetail", c)
}

func DeleteDocDetail(c *gin.Context) {
	fmt.Println("delete docDetail", c)
}
