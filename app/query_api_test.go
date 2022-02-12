package main

import (
	"testing"
)

func TestQueryConstructor(t *testing.T) {
	query := ApiQuery{}
	response := ApiResponse{}

	err := AddStringFilter(&query, "launchpad", "5e9e4501f509094ba4566f84")
	if err != nil {
		t.Error(err)
	}

	err = AddBoolFilter(&query, "upcoming", true)
	if err != nil {
		t.Error(err)
	}

	err = AddBoolFilter(&query, "tbd", false)
	if err != nil {
		t.Error(err)
	}

	err = AddDateFilter(&query, "2022-02-20")
	if err != nil {
		t.Error(err)
	}

	err = SendApiQuery(LaunchesEndpointV5, &query, &response)
	if err != nil {
		t.Error(err)
	}
}
