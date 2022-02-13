package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

var dbUrl = os.Getenv("DATABASE_URL")
var dbDriver = os.Getenv("DATABASE_DRIVER")

func applyLogSettings() {
	// stderr -> stdout for convenience
	log.SetOutput(os.Stdout)
}

func applyDbMigrations() {
	// Create Booking relation
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
}

func assignRoutes() *gin.Engine {
	// Add POST/GET/DELETE endpoints for Booking resource
	router := gin.Default()
	router.GET("/booking", func(c *gin.Context) {
		bookings := []Booking{}
		res := dbHandle.Find(&bookings)
		if res.Error != nil {
			log.Printf("Failed to retrieve bookings: %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, bookings)
	})

	router.POST("/booking", func(c *gin.Context) {
		log.Printf("Create booking")

		var bookingRequest BookingCreateRequest
		if err := c.ShouldBindJSON(&bookingRequest); err != nil {
			log.Printf("Failed to parse incoming booking request: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		booking, err := CreateBooking(&bookingRequest)
		if err != nil {
			log.Printf("Failed to validate booking: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}

		res := dbHandle.Create(booking)
		if res.Error != nil {
			log.Printf("Failed to create booking %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		c.JSON(http.StatusOK, booking)
	})

	router.DELETE("/booking/:id", func(c *gin.Context) {
		bookingId := c.Param("id")
		log.Printf("Delete booking: %v", bookingId)

		u, err := strconv.ParseUint(bookingId, 10, 64)
		if err != nil {
			log.Printf("Failed to parse booking id: %v", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
		booking := Booking{ID: u}

		res := dbHandle.Delete(&booking)
		if res.Error != nil {
			log.Printf("Failed to delete booking: %v", res.Error)
			c.JSON(http.StatusBadRequest, gin.H{"error": res.Error.Error()})
			return
		}

		log.Printf("Deleted bookings: %v", res.RowsAffected)
		c.JSON(http.StatusOK, gin.H{})
	})

	return router
}

func main() {
	applyLogSettings()
	applyDbMigrations()
	assignRoutes().Run() // 0.0.0.0:8080
}
