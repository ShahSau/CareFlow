package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetUsers(c *gin.Context) {

	fmt.Println("get all users", c)

	c.JSON(200, gin.H{
		"message": "get all users",
	})
}

func GetUser(c *gin.Context) {
	fmt.Println("get user", c)
}

func CreateUser(c *gin.Context) {
	fmt.Println("create user", c)
}

func UpdateUser(c *gin.Context) {
	fmt.Println("update user", c)
}

func DeleteUser(c *gin.Context) {
	fmt.Println("delete user", c)
}
