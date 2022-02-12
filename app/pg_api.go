package main

import (
	"database/sql"
	"errors"
	"log"

	_ "github.com/lib/pq"
)

var dbHandle *sql.DB

func GetDbHandle(dbDriver string, dbUrl string) (*sql.DB, error) {
	if dbHandle == nil {
		newHandle, err := sql.Open(dbDriver, dbUrl)
		if err != nil {
			return nil, err
		}
		dbHandle = newHandle
	}

	return dbHandle, nil
}

func FreeDbHandle(dbHandle *sql.DB) error {
	if dbHandle == nil {
		return errors.New("Can't free uninitialized DB handle")
	}
	return dbHandle.Close()
}

func RetrieveBookings(dbHandle *sql.DB, bookings []Booking) error {
	if dbHandle == nil {
		return errors.New("Can't perform DB query, DB handle uninitialized")
	}

	rows, err := dbHandle.Query("select * from bookings")
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		booking := Booking{}
		err := rows.Scan(&booking)
		if err != nil {
			return err
		}
		bookings = append(bookings, booking)
		if err == nil {
			log.Fatal(booking)
		}
	}
	err = rows.Err()
	if err != nil {
		return err
	}
	return nil
}
