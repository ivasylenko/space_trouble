package main

import (
	"bytes"
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"time"
)

var LaunchesEndpointV5 = "https://api.spacexdata.com/v5/launches/query"
var dateLayout = "2006-01-02"

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
		return err
	}

	requestBody := bytes.NewBuffer(jsonBody)

	resp, err := http.Post(endpoint, "application/json", requestBody)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	responseBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}

	err = json.Unmarshal(responseBody, &response)
	if err != nil {
		return err
	}

	return nil
}

func AddDateFilter(query *ApiQuery, date string) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query["date_utc"]; ok {
		return errors.New("date_utc was already set")
	}

	time_from, err := time.Parse(dateLayout, date)
	if err != nil {
		return err
	}

	time_to := time_from.Add(time.Hour * 24)

	query.Query["date_utc"] = map[string]string{
		"$gte": time_from.Format(time.RFC3339),
		"$lte": time_to.Format(time.RFC3339),
	}
	return nil
}

func AddStringFilter(query *ApiQuery, key string, value string) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query[key]; ok {
		return errors.New("key was already set")
	}
	query.Query[key] = value
	return nil
}

func AddBoolFilter(query *ApiQuery, key string, value bool) error {
	if query.Query == nil {
		query.Query = map[string]interface{}{}
	}
	if _, ok := query.Query[key]; ok {
		return errors.New("key was already set")
	}
	query.Query[key] = value
	return nil
}
