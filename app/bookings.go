package main

import (
	"encoding/json"
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

type DateTime time.Time

type Booking struct {
	ID            uint `gorm:"primaryKey"`
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

type BookingDeleteRequestById struct {
	ID uint `gorm:"primaryKey" json:"id" binding:"required"`
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

func FromBookingRequest(bookingRequest *BookingCreateRequest) *Booking {
	return &Booking{
		FirstName:     bookingRequest.FirstName,
		LastName:      bookingRequest.LastName,
		Gender:        bookingRequest.Gender,
		Birthday:      time.Time(bookingRequest.Birthday),
		LaunchpadID:   bookingRequest.LaunchpadID,
		DestinationID: bookingRequest.DestinationID,
		LaunchDate:    time.Time(bookingRequest.LaunchDate),
	}
}
