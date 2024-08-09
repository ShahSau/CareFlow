package controllers

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/ShahSau/CareFlow/backend/database"
	"github.com/ShahSau/CareFlow/backend/helpers"
	"github.com/ShahSau/CareFlow/backend/models"
	"github.com/ShahSau/CareFlow/backend/types"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var userCollection *mongo.Collection = database.GetCollection(database.DB, "users")

func HashPassword(password string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	if err != nil {
		log.Panic("Error hashing password")
	}
	return string(bytes)
}

func ComparePassword(hashedPassword string, password string) (bool, string) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
	if err != nil {
		return false, "Invalid password"
	}
	return true, "Password is valid"
}

// use only login function, (if new user use Login functionalities, if new user use Register functionalities)
func Login(c *gin.Context) {

	fmt.Println("login", c)

	c.JSON(200, gin.H{
		"message": "login",
	})
}

func Register(c *gin.Context) {
	var user types.RegisterUser
	var foundUser models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	// checking if user already exists
	count, err := userCollection.CountDocuments(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}})
	if err != nil {
		c.JSON(500, gin.H{
			"error": "Error while checking user",
		})
		return
	}

	//Email already exists
	if count > 0 {
		err := userCollection.FindOne(c.Request.Context(), bson.D{{Key: "email", Value: user.Email}}).Decode(&foundUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		passwordValid, msg := ComparePassword(foundUser.Password, user.Password)

		if !passwordValid {
			c.JSON(http.StatusBadRequest, gin.H{"error": true, "message": msg, "status": http.StatusBadRequest, "success": false})
			return
		}

		token, refreshToken, _ := helpers.GenerateAllTokens(foundUser.User_id, foundUser.Email, foundUser.Name, foundUser.Phone)

		helpers.UpdateAllTokens(token, refreshToken, foundUser.User_id)

		c.JSON(http.StatusOK, gin.H{"error": false, "message": "User logged in successfully", "data": foundUser, "status": http.StatusOK, "success": true, "token": token, "refresh_token": refreshToken, "new_user": false})
	}
	if count == 0 {
		var newUser models.User

		// Inserting user to database
		newUser.Name = user.Name
		newUser.Email = user.Email
		newUser.Phone = user.Phone
		newUser.Role = user.Role
		password := HashPassword(user.Password)
		newUser.Password = password
		newUser.CreatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newUser.UpdatedAt, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
		newUser.ID = primitive.NewObjectID()
		newUser.User_id = newUser.ID.Hex()

		token, refreshToken, _ := helpers.GenerateAllTokens(newUser.User_id, newUser.Email, newUser.Name, newUser.Phone)
		newUser.Token = token
		newUser.RefreshToken = refreshToken

		_, err = userCollection.InsertOne(c.Request.Context(), newUser)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusCreated, gin.H{"error": false, "message": "User created successfully", "data": newUser, "status": http.StatusCreated, "success": true, "token": token, "refresh_token": refreshToken, "new_user": true})
	}
}

func Logout(c *gin.Context) {
	fmt.Println("logout", c)
}

func Test(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "test",
		"status":  http.StatusNoContent,
		"success": true,
		"error":   false,
		"data": map[string]string{
			"test":  "test",
			"test2": "test2",
		},
	})
}
