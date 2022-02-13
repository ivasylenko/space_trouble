package main

import (
	"log"
	"testing"
	"time"
)

func TestQueryConstructor(t *testing.T) {
	launchDate, _ := time.Parse("2006-01-02", "2022-02-20")
	query := ApiQuery{
		Query: map[string]interface{}{
			"launchpad": "5e9e4501f509094ba4566f84",
			"upcoming":  true,
			"tbd":       false,
			"date_utc": map[string]string{
				"$gte": launchDate.Format(time.RFC3339),
				"$lte": launchDate.Add(time.Hour * 24).Format(time.RFC3339),
			},
		},
	}
	response := ApiResponse{}

	err := SendApiQuery(LaunchesEndpointV5, &query, &response)
	if err != nil {
		t.Error(err)
	}
	log.Printf("Response: %v", response.TotalDocs)
}
