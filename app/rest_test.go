package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"testing"
	"time"
)

func TestPostBooking(t *testing.T) {
	// booking := Booking{LastName: "LALALALA",
	// 	FirstName: "BLABLABLA", Gender: "Blsb",
	// 	Birthday: "1993-01-15", LaunchpadID: "launch-id",
	// 	DestinationID: "destination-id", LaunchDate: "2022-11-21",
	// }
	booking_request := map[string]string{
		"first_name":     "Baba",
		"last_name":      "O'riley",
		"gender":         "Male",
		"birthday":       time.Now().Format("2006-01-02"),
		"launchpad_id":   "lllll-ididididi",
		"destination_id": "ddddd-idididid",
		"launch_date":    "1996-01-15",
	}

	jsonBody, err := json.Marshal(&booking_request)
	if err != nil {
		t.Errorf("unexpected error while marshling: %v", err)
	}

	requestBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post("http://127.0.0.1:8080/booking", "application/json", requestBody)
	if err != nil {
		t.Errorf("unexpected error response: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error readinf response: %v", err)
	}
	t.Errorf("unexpected error unmarshaling response: %v", string(responseBody))
}

func TestRetrieveBookings(t *testing.T) {
	// bookings := []Booking{}

	resp, err := http.Get("http://127.0.0.1:8080/booking")
	if err != nil {
		t.Errorf("unexpected error response: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error readinf response: %v", err)
	}
	t.Errorf("unexpected error unmarshaling response: %v", string(responseBody))
}
