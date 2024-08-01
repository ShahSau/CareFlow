package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {

	fmt.Println("login", c)

	c.JSON(200, gin.H{
		"message": "login",
	})
}

func Register(c *gin.Context) {
	fmt.Println("register", c)
}

func Logout(c *gin.Context) {
	fmt.Println("logout", c)
}
