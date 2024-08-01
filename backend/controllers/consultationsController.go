package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetConsultations(c *gin.Context) {

	fmt.Println("get all consultations", c)

	c.JSON(200, gin.H{
		"message": "get all consultations",
	})
}

func GetConsultation(c *gin.Context) {
	fmt.Println("get consultation", c)
}

func CreateConsultation(c *gin.Context) {
	fmt.Println("create consultation", c)
}

func UpdateConsultation(c *gin.Context) {
	fmt.Println("update consultation", c)
}

func DeleteConsultation(c *gin.Context) {
	fmt.Println("delete consultation", c)
}

func GetConsultationByPatient(c *gin.Context) {
	fmt.Println("get consultation by patient", c)
}

func GetConsultationByDoctor(c *gin.Context) {
	fmt.Println("get consultation by doctor", c)
}

func GetConsultationByDate(c *gin.Context) {
	fmt.Println("get consultation by date", c)
}
