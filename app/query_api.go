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

func SendApiQuery(endpoint string, query *ApiQuery, response *ApiResponse) error {
	jsonBody, err := json.Marshal(query)
	if err != nil {
		return fmt.Errorf("failed to format query request: %v, %v", query, err)
	}

	requestBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(endpoint, "application/json", requestBody)
	if err != nil {
		return fmt.Errorf("failed to send query: %v, %v", string(jsonBody), err)
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read api response: %v", err)
	}

	log.Printf("Received API response: %v", string(responseBody))

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return fmt.Errorf("failed to unmarshal api response: %v", string(responseBody))
	}

	return nil
}

func AddDateFilter(query *ApiQuery, date time.Time) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query["date_utc"]; ok {
		return fmt.Errorf("failed to date filter: %v as it is already set", date)
	}

	date_to := date.Add(time.Hour * 24)

	query.Query["date_utc"] = map[string]string{
		"$gte": date.Format(time.RFC3339),
		"$lte": date_to.Format(time.RFC3339),
	}
	return nil
}

func AddStringFilter(query *ApiQuery, key string, value string) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query[key]; ok {
		return fmt.Errorf("failed to add string filter {%v: %v}", key, value)
	}
	query.Query[key] = value
	return nil
}

func AddBoolFilter(query *ApiQuery, key string, value bool) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query[key]; ok {
		return fmt.Errorf("failed to add bool filter {%v: %v}", key, value)
	}
	query.Query[key] = value
	return nil
}
