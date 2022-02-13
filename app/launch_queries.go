package main

import (
	"log"
	"time"
)

func CheckLaunchpadAvailable(launchpad_id string, launch_date time.Time) (bool, error) {
	//Query if there is Launch on given `launchpad_id` and date `launch_date
	//return if avaiable for launch
	query := ApiQuery{}
	response := ApiResponse{}

	err := AddStringFilter(&query, "launchpad", launchpad_id)
	if err != nil {
		return false, err
	}

	err = AddBoolFilter(&query, "upcoming", true)
	if err != nil {
		return false, err
	}

	err = AddBoolFilter(&query, "tbd", false)
	if err != nil {
		return false, err
	}

	err = AddDateFilter(&query, launch_date)
	if err != nil {
		return false, err
	}

	err = SendApiQuery(LaunchesEndpointV5, &query, &response)
	if err != nil {
		return false, err
	}

	if response.TotalDocs > 0 {
		log.Printf("Launchpad: %s is not avaiable on: %s", launchpad_id, launch_date)
		return false, nil
	}

	return true, nil
}
