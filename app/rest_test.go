package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
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
		"launchpad_id":   "5e9e4501f509094ba4566f84",
		"destination_id": "Ganymede",
		"launch_date":    "2022-02-20",
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
	log.Printf("Response: %v", string(responseBody))
}

func TestGetBookings(t *testing.T) {
	resp, err := http.Get("http://127.0.0.1:8080/booking")
	if err != nil {
		t.Errorf("unexpected error response: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("unexpected error reading response: %v", err)
	}
	log.Printf("Response: %v", string(responseBody))
}

func TestDeleteBooking(t *testing.T) {

	client := &http.Client{}

	req, err := http.NewRequest("DELETE", "http://127.0.0.1:8080/booking/8", nil)
	if err != nil {
		t.Errorf("failed to construct request: %v", err)
	}

	// Fetch Request
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("error sending delete request : %v", err)
		return
	}
	defer resp.Body.Close()

	// Read Response Body
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		t.Errorf("error reading response : %v", err)
		return
	}

	log.Printf("response Status : %v", resp.Status)
	log.Printf("response Headers : %v", resp.Header)
	log.Printf("response Body : %v", string(respBody))
}
