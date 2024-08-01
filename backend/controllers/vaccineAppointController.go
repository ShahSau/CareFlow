package controllers

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func GetVaccineAppointments(c *gin.Context) {

	fmt.Println("get all vaccine appointments", c)

	c.JSON(200, gin.H{
		"message": "get all vaccine appointments",
	})
}

func GetVaccineAppointment(c *gin.Context) {
	fmt.Println("get vaccine appointment", c)
}

func CreateVaccineAppointment(c *gin.Context) {
	fmt.Println("create vaccine appointment", c)
}

func UpdateVaccineAppointment(c *gin.Context) {
	fmt.Println("update vaccine appointment", c)
}

func DeleteVaccineAppointment(c *gin.Context) {
	fmt.Println("delete vaccine appointment", c)
}

func GetVaccineAppointmentsByVaccine(c *gin.Context) {
	fmt.Println("get vaccine appointments by vaccine", c)
}

func GetVaccineAppointmentsByUser(c *gin.Context) {
	fmt.Println("get vaccine appointments by user", c)
}

func GetVaccineAppointmentsByDate(c *gin.Context) {
	fmt.Println("get vaccine appointments by date", c)
}
