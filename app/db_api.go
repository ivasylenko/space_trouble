package main

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var dbHandle *gorm.DB

func GetDbHandle(dbDriver string, dbUrl string) (*gorm.DB, error) {
	if dbHandle == nil {
		sqlDB, err := sql.Open(dbDriver, dbUrl)
		if err != nil {
			return nil, err
		}

		dbHandle, err = gorm.Open(postgres.New(postgres.Config{
			Conn: sqlDB,
		}), &gorm.Config{})
		if err != nil {
			return nil, err
		}
	}

	return dbHandle, nil
}
