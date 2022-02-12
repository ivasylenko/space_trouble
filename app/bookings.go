package main

import (
	"fmt"
	"time"
)

type GenderName string
type DestinationName string

const (
	Male   GenderName = "Male"
	Female GenderName = "Female"
)

const (
	Mars     DestinationName = "Mars"
	Moon     DestinationName = "Moon"
	Pluto    DestinationName = "Pluto"
	Asteroid DestinationName = "Asteroid"
	Belt     DestinationName = "Belt"
	Europa   DestinationName = "Europa"
	Titan    DestinationName = "Titan"
	Ganymede DestinationName = "Ganymede"
)

type BookingDetails struct {
	FirstName     string          `json:"first_name"`
	LastName      string          `json:"last_name"`
	Gender        GenderName      `json:"gender"`
	Birthday      time.Time       `json:"birthday"`
	LaunchpadID   string          `json:"launchpad_id"`
	DestinationID DestinationName `json:"destination_id"`
	LaunchDate    time.Time       `json:"launch_date"`
}

type Booking struct {
	BookingDetails
	BookingID string `json:"booking_id"`
}

type DeleteBookingRequest struct {
	Booking
}

func CreateBooking(booking_request *BookingDetails) (string, error) {
	fmt.Printf("%+v\n", booking_request)
	return "booking-id", nil
}

func DeleteBooking(booking_request *DeleteBookingRequest) error {
	return nil
}
