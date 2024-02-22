package db

import (
	"log"
	"telmed_backend/config"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB = nil

func GetDB() *gorm.DB {
	if db != nil {
		return db
	}
	db, err := Connect()
	if err != nil {
		panic("Some error occured while conencting to db")
	}

	return db

}

func Connect() (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(config.DB_URI), &gorm.Config{})
	if err != nil {
		log.Println("Error coonection to db", err)
		return nil, err
	}
	log.Println("Connected to cockroach")

	return db, nil
}
