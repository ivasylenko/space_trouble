package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/booking", func(c *gin.Context) {
		bookings, err := GetBookings()
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, bookings)
	})

	r.POST("/booking", func(c *gin.Context) {
		var booking_request BookingDetails
		if err := c.ShouldBindJSON(&booking_request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		booking_id, err := CreateBooking(&booking_request)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"booking_id": booking_id})
	})

	r.DELETE("/booking", func(c *gin.Context) {
		var booking_request DeleteBookingRequest
		if err := c.ShouldBindJSON(&booking_request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := DeleteBooking(&booking_request); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, gin.H{"deleted": "deleted"})
	})

	r.Run() // 0.0.0.0:8080
}
