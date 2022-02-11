package main

import (
	"testing"
)

func TestLaunchpadAPI(t *testing.T) {
	SendApiQuery(LaunchesEndpointV5, &ApiQuery{
		Query: map[string]string{
			"launchpad": "5e9e4501f509094ba4566f84",
		},
	})
}

func TestDateToQueryTerm(t *testing.T) {
	// dateOFLaunch := time.Now().UTC().Format(time.RFC3339)
	dateOFLaunch := "2022-01-22"
	result, err := DateFilterTerm(dateOFLaunch)
	if err != nil {
		t.Fatal(err)
	}
	t.Fatal((result))
}
