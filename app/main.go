package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var dbUrl = os.Getenv("DATABASE_URL")
var dbDriver = os.Getenv("DATABASE_DRIVER")

func main() {
	dbHandle, err := GetDbHandle(dbDriver, dbUrl)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = dbHandle.AutoMigrate(&Booking{})
	if err != nil {
		fmt.Println(err)
		return
	}

	route := gin.Default()

	route.GET("/booking", func(c *gin.Context) {
		bookings := []Booking{}
		res := dbHandle.Find(&bookings)
		if res.Error != nil {
			log.Printf("Failed to retrieve bookings: %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, bookings)
	})

	route.POST("/booking", func(c *gin.Context) {
		var bookingRequest BookingCreateRequest
		if err := c.ShouldBindJSON(&bookingRequest); err != nil {
			log.Printf("Failed to parse incoming booking request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		booking := FromBookingRequest(&bookingRequest)
		res := dbHandle.Create(booking)
		if res.Error != nil {
			log.Printf("Failed to create booking %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, booking)
	})

	route.DELETE("/booking", func(c *gin.Context) {
		var bookingRequest BookingDeleteRequestById
		if err := c.ShouldBindJSON(&bookingRequest); err != nil {
			log.Printf("Failed to parse incoming booking request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		res := dbHandle.Delete(&bookingRequest)
		if res.Error != nil {
			log.Printf("Failed to delete booking: %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, bookingRequest)
	})

	route.Run() // 0.0.0.0:8080
}
