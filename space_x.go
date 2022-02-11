package main

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

var LaunchpadsEndpointV4 = "https://api.spacexdata.com/v4/launchpads/query"
var LaunchesEndpointV5 = "https://api.spacexdata.com/v5/launches/query"
var dateLayout = "2006-01-02"

type ApiQuery struct {
	Query   map[string]string `json:"query,omitempty"`
	Options map[string]string `default:"{}" json:"options,omitempty"`
}

func SendApiQuery(endpoint string, query *ApiQuery) {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		log.Fatalln(err)
		return
	}

	requestBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(endpoint, "application/json", requestBody)
	if err != nil {
		log.Fatalf("An Error Occured %v", err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	sb := string(responseBody)
	log.Printf(sb)
}

func DateFilterTerm(launch_date string) (map[string]string, error) {
	time_from, err := time.Parse(dateLayout, launch_date)
	if err != nil {
		log.Fatalf("Failed to parsing date of launch - %v", err)
		return nil, nil
	}

	time_to := time_from.Add(time.Hour * 24)

	return map[string]string{
		"$gte": time_from.Format(time.RFC3339),
		"$lte": time_to.Format(time.RFC3339),
	}, nil
}
