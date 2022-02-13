package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var LaunchesEndpointV5 = "https://api.spacexdata.com/v5/launches/query"

type ApiQuery struct {
	Query   map[string]interface{} `json:"query,omitempty"`
	Options map[string]interface{} `default:"{}" json:"options,omitempty"`
}

type ApiResponse struct {
	TotalDocs int           `json:"totalDocs"`
	Docs      []interface{} `json:"docs"`
}

func CheckLaunchpadAvailable(launchpadId string, launchDate time.Time) (bool, error) {
	//Query if there is Launch on given `launchpadId` and date `launchDate`
	//return if avaiable for launch

	response := ApiResponse{}
	query := ApiQuery{
		Query: map[string]interface{}{
			"launchpad": launchpadId,
			"upcoming":  true,
			"tbd":       false,
			"date_utc": map[string]string{
				"$gte": launchDate.Format(time.RFC3339),
				"$lte": launchDate.Add(time.Hour * 24).Format(time.RFC3339),
			},
		},
	}

	err := SendApiQuery(LaunchesEndpointV5, &query, &response)
	if err != nil {
		return false, err
	}

	// Check if there are launches from given launchpad on the date
	if response.TotalDocs > 0 {
		log.Printf("launchpad: %v is not avaiable on: %v", launchpadId, launchDate)
		return false, nil
	}

	return true, nil
}

func SendApiQuery(endpoint string, query *ApiQuery, response *ApiResponse) error {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("can't marshal query: %v", err)
	}

	requestBody := bytes.NewBuffer(jsonBody)

	log.Printf("Send api query: %v", string(jsonBody))

	resp, err := http.Post(endpoint, "application/json", requestBody)
	if err != nil {
		return fmt.Errorf("can't send query: %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("can't read api response: %v", err)
	}

	log.Printf("Received api response: %v", string(responseBody))

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		log.Printf("Can't unmarshal API response: %v - %v", string(responseBody), err)
		return fmt.Errorf("can't unmarshal api response: %v", err)
	}

	return nil
}
