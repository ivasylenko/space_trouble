package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"testing"
	"time"
)

func TestPostBookingBusy(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking"
	err := runPostTest(bookingUrl,
		&map[string]string{
			"first_name":     "Pamela",
			"last_name":      "Mars",
			"gender":         "Male",
			"birthday":       time.Now().Format("2006-01-02"),
			"launchpad_id":   "5e9e4501f509094ba4566f84",
			"destination_id": "Ganymede",
			"launch_date":    "2022-02-20",
		},
		http.StatusBadRequest)
	if err != nil {
		t.Error(err)
	}
}

func TestPostBookingWrongDestination(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking"
	err := runPostTest(bookingUrl,
		&map[string]string{
			"first_name":     "Pamela",
			"last_name":      "Mars",
			"gender":         "Male",
			"birthday":       time.Now().Format("2006-01-02"),
			"launchpad_id":   "5e9e4501f509094ba4566f84",
			"destination_id": "Neptun",
			"launch_date":    "2022-02-20",
		},
		http.StatusBadRequest)
	if err != nil {
		t.Error(err)
	}
}

func TestPostDestinationIsNotScheduledOnDay(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking"
	err := runPostTest(bookingUrl,
		&map[string]string{
			"first_name":     "Pamela",
			"last_name":      "Mars",
			"gender":         "Male",
			"birthday":       time.Now().Format("2006-01-02"),
			"launchpad_id":   "5e9e4501f509094ba4566f84",
			"destination_id": "Ganymede",
			"launch_date":    "2022-02-21",
		},
		http.StatusBadRequest)
	if err != nil {
		t.Error(err)
	}
}

func TestPostSuccess(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking"
	err := runPostTest(bookingUrl,
		&map[string]string{
			"first_name":     "Pamela",
			"last_name":      "Mars",
			"gender":         "Male",
			"birthday":       time.Now().Format("2006-01-02"),
			"launchpad_id":   "5e9e4501f509094ba4566f84",
			"destination_id": "Mars",
			"launch_date":    "2022-02-21",
		},
		http.StatusOK)
	if err != nil {
		t.Error(err)
	}
}

func TestGetBookings(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking"
	err := runGetTest(bookingUrl, 2)
	if err != nil {
		t.Error(err)
	}
}

func TestDeleteBooking(t *testing.T) {
	bookingUrl := "http://127.0.0.1:8080/booking/11"
	err := runDeleteTest(bookingUrl)
	if err != nil {
		t.Error(err)
	}
}

func runDeleteTest(bookingUrl string) error {
	client := &http.Client{}

	req, err := http.NewRequest("DELETE", bookingUrl, nil)
	if err != nil {
		return fmt.Errorf("failed to construct request: %v", err)
	}

	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("error sending delete request : %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("error reading response : %v", err)
	}

	log.Printf("response Status : %v", resp.Status)
	log.Printf("response Headers : %v", resp.Header)
	log.Printf("response Body : %v", string(responseBody))
	return nil
}

func runGetTest(bookingUrl string, expectedItems int) error {
	resp, err := http.Get(bookingUrl)
	if err != nil {
		return fmt.Errorf("unexpected error response: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unexpected error reading response: %v", err)
	}

	log.Printf("response Status : %v", resp.Status)
	log.Printf("response Headers : %v", resp.Header)
	log.Printf("response Body : %v", string(responseBody))

	var bookings []interface{}
	err = json.Unmarshal(responseBody, &bookings)
	if err != nil {
		return fmt.Errorf("unexpected error while marshling: %v", err)
	}

	if len(bookings) != expectedItems {
		return fmt.Errorf("expected: %v bookings, found: %v", expectedItems, len(bookings))
	}
	return nil
}

func runPostTest(bookingUrl string, request *map[string]string, expected_code int) error {
	jsonBody, err := json.Marshal(request)
	if err != nil {
		return fmt.Errorf("unexpected error while marshling: %v", err)
	}

	requestBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(bookingUrl, "application/json", requestBody)
	if err != nil {
		return fmt.Errorf("unexpected error response: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("unexpected error readinf response: %v", err)
	}

	log.Printf("response Status : %v", resp.Status)
	log.Printf("response Headers : %v", resp.Header)
	log.Printf("response Body : %v", string(responseBody))

	if resp.StatusCode != expected_code {
		return fmt.Errorf("wrong response: %v", resp.StatusCode)
	}

	return nil
}
