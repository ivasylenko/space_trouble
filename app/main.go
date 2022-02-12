package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

var dbUrl = os.Getenv("DATABASE_URL")
var dbDriver = "postgres"

func main() {
	r := gin.Default()

	r.GET("/booking", func(c *gin.Context) {
		dbHandle, err := GetDbHandle(dbDriver, dbUrl)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		bookings := []Booking{}

		err = RetrieveBookings(dbHandle, bookings)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, bookings)
	})

	r.POST("/booking", func(c *gin.Context) {
		var booking_request BookingDetails
		if err := c.ShouldBindJSON(&booking_request); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		booking_id, err := CreateBooking(&booking_request)
		if err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"booking_id": booking_id})
	})

	r.DELETE("/booking", func(c *gin.Context) {
		var booking_request DeleteBookingRequest
		if err := c.ShouldBindJSON(&booking_request); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := DeleteBooking(&booking_request); err != nil {
			fmt.Println(err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"deleted": "deleted"})
	})

	r.Run() // 0.0.0.0:8080
}
