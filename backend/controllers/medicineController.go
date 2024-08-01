package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetMedicines(c *gin.Context) {

	fmt.Println("get all medicines", c)

	c.JSON(200, gin.H{
		"message": "get all medicines",
	})
}

func GetMedicine(c *gin.Context) {
	fmt.Println("get medicine", c)
}

func CreateMedicine(c *gin.Context) {
	fmt.Println("create medicine", c)
}

func UpdateMedicine(c *gin.Context) {
	fmt.Println("update medicine", c)
}

func DeleteMedicine(c *gin.Context) {
	fmt.Println("delete medicine", c)
}

func GetMedicineByDisease(c *gin.Context) {
	fmt.Println("get medicine by disease", c)
}

func GetMedicineBySpecialization(c *gin.Context) {
	fmt.Println("get medicine by specialization", c)
}

func GetMedicineByPrescription(c *gin.Context) {
	fmt.Println("get medicine by prescription", c)
}

func GetMedicineByPatient(c *gin.Context) {
	fmt.Println("get medicine by patient", c)
}
