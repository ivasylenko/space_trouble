package main

import (
	"database/sql"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var dbUrl = os.Getenv("DATABASE_URL")
var dbDriver = "postgres"

func RetrieveBookings(bookings []Booking) error {
	db, err := sql.Open(dbDriver, dbUrl)
	if err != nil {
		return err
	}

	rows, err := db.Query("select * from bookings")
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
