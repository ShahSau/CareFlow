package controllers

import (
	"fmt"
	"net/http"

	"github.com/ShahSau/CareFlow/backend/models"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
)

func GetUsers(c *gin.Context) {

	fmt.Println("get all users", c)

	c.JSON(200, gin.H{
		"message": "get all users",
	})
}

func GetUser(c *gin.Context) {
	userId := c.Param("id")

	var user models.User

	defer c.Request.Body.Close()

	err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "user_id", Value: userId}}).Decode(&user)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user, "status": http.StatusOK, "success": true, "error": false, "message": "User retrieved successfully"})

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
