package main

import (
	"encoding/json"
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
	Mars         DestinationName = "Mars"
	Moon         DestinationName = "Moon"
	Pluto        DestinationName = "Pluto"
	AsteroidBelt DestinationName = "Asteroid Belt"
	Europa       DestinationName = "Europa"
	Titan        DestinationName = "Titan"
	Ganymede     DestinationName = "Ganymede"
)

type DateTime time.Time

var schedule = map[time.Weekday]DestinationName{
	time.Monday:    Mars,
	time.Tuesday:   Moon,
	time.Wednesday: Pluto,
	time.Thursday:  AsteroidBelt,
	time.Friday:    Europa,
	time.Saturday:  Titan,
	time.Sunday:    Ganymede,
}

var validDestinations = map[DestinationName]bool{
	Mars: true, Moon: true, Pluto: true, AsteroidBelt: true, Europa: true, Titan: true, Ganymede: true,
}

type Booking struct {
	ID            uint64 `gorm:"primaryKey"`
	FirstName     string
	LastName      string
	Gender        GenderName
	Birthday      time.Time
	LaunchpadID   string
	DestinationID DestinationName
	LaunchDate    time.Time
}

type BookingCreateRequest struct {
	FirstName     string          `json:"first_name" binding:"required"`
	LastName      string          `json:"last_name" binding:"required"`
	Gender        GenderName      `json:"gender" binding:"required"`
	Birthday      DateTime        `json:"birthday" binding:"required"`
	LaunchpadID   string          `json:"launchpad_id" binding:"required"`
	DestinationID DestinationName `json:"destination_id" binding:"required"`
	LaunchDate    DateTime        `json:"launch_date" binding:"required"`
}

func (mt *DateTime) UnmarshalJSON(bs []byte) error {
	var s string
	err := json.Unmarshal(bs, &s)
	if err != nil {
		return err
	}
	t, err := time.ParseInLocation("2006-01-02", s, time.UTC)
	if err != nil {
		return err
	}
	*mt = DateTime(t)
	return nil
}

func CreateBooking(bookingRequest *BookingCreateRequest) (*Booking, error) {
	// Validate Booking request and create Booking

	destinationID := bookingRequest.DestinationID
	if _, ok := validDestinations[destinationID]; !ok {
		return nil, fmt.Errorf("unknown destination: %v", destinationID)
	}

	launchDate := time.Time(bookingRequest.LaunchDate).UTC()
	if v, ok := schedule[launchDate.Weekday()]; ok {
		if v != bookingRequest.DestinationID {
			return nil, fmt.Errorf("no flights to: %v on: %v", destinationID, launchDate.Weekday())
		}
	}

	launchpadID := bookingRequest.LaunchpadID
	ok, err := CheckLaunchpadAvailable(launchpadID, launchDate)
	if err != nil {
		return nil, err
	}

	if !ok {
		return nil, fmt.Errorf("date: %v is not avaiable for launchpad: %v", launchDate, launchpadID)
	}

	return &Booking{
		FirstName:     bookingRequest.FirstName,
		LastName:      bookingRequest.LastName,
		Gender:        bookingRequest.Gender,
		Birthday:      time.Time(bookingRequest.Birthday),
		LaunchpadID:   launchpadID,
		DestinationID: destinationID,
		LaunchDate:    launchDate,
	}, nil
}
